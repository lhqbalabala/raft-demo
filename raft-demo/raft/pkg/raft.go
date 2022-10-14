package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"raft/models"
	"raft/pkg/utils"
	"sync"
	"time"
)

type RaftState int

const (
	Follower RaftState = iota
	Candidate
	Leader
	Dead
	New
)

type NodeInfo struct {
	ID      string
	Address string
	Port    string
}

// 声明raft节点类型
type Raft struct {
	node *NodeInfo
	//本节点获得的投票数
	vote int
	//线程锁
	lock sync.Mutex
	//节点编号
	id string

	//当前任期
	currentTerm int32
	//为哪个节点投票
	votedFor string
	//当前节点的log
	log []*models.LogEntry

	//当前节点状态
	//0 follower  1 candidate  2 leader 3 dead
	state RaftState

	//超时重置时间：
	//1、收到心跳信号
	//2、投票给了他人
	//3、变为候选人
	electionResetEvent time.Time

	//当前节点的领导
	currentLeader string
	//心跳超时时间(单位：秒)
	timeoutDuration time.Duration

	//peerIds 列出了所有id ID。
	peerIds []string

	commitIndex int32 //已知已提交的最高的日志条目的索引（初始值为-1，单调递增）
	lastApplied int32 //已经被应用到状态机的最高的日志条目的索引（初始值为-1，单调递增）

	nextIndex  map[string]int32 //对于每一台服务器，发送到该服务器的下一个日志条目的索引（初始值为领导人最后的日志条目的索引+1）
	matchIndex map[string]int32 //对于每一台服务器，已知的已经发送到该服务器的最高日志条目的索引（初始值为0，单调递增）
	server     *Server

	//可以提交的信号
	SignToCommit chan struct{}
	//提交数据通道
	CommitChannel chan models.LogEntry
	//向从发送数据的信号
	AppendToFollowerChannel chan struct{}

	//起始位置
	StartIndex int32
	//命令ID
	CommondIds map[int32]bool
	Stroage    *KvStorage

	Snap     *Snap
	Snapbool map[string]bool
}
type Snap struct {
	Stroage *KvStorage
	Index   int32
}

func (s RaftState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	case New:
		return "New"
	default:
		panic("unreachable")
	}
}

func NewRaft(id, address, port string, peerIds []string, startchan chan struct{}) *Raft {
	node := &NodeInfo{id, address, port}
	rf := new(Raft)
	rf.CommondIds = make(map[int32]bool)
	rf.Stroage = &KvStorage{}
	rf.peerIds = peerIds
	rf.Snapbool = make(map[string]bool)
	rf.Snap = &Snap{}
	//节点信息
	rf.node = node
	//当前节点获得票数
	rf.setVote(0)
	//编号
	rf.id = id
	//初始未投票
	rf.setVoteFor("-1")
	//初始为New
	rf.setStatus(New)
	rf.SignToCommit = make(chan struct{}, 16)
	rf.CommitChannel = make(chan models.LogEntry, 1024)
	rf.AppendToFollowerChannel = make(chan struct{}, 16)
	//最初没有领导
	rf.setCurrentLeader("-1")
	//设置任期
	rf.setTerm(0)
	rf.commitIndex = -1
	rf.lastApplied = -1
	rf.nextIndex = make(map[string]int32)
	rf.matchIndex = make(map[string]int32)
	rf.StartIndex = 0
	//rf.load()
	go func() {
	_:
		<-startchan
		rf.electionResetEvent = time.Now()
		go rf.TocommitTomachine()
		go rf.runElectionTimer()
		go rf.CommitToMachine()
		//开启后台线程随机删除一些过期的key
		go rf.Stroage.runercleanTimer()
		//go rf.AutoFlush()
	}()

	return rf
}
func (rf *Raft) load() {
	var Term int32
	var VoteFor string
	var Log []*models.LogEntry
	var CommitIndex int32
	var LastApplied int32
	utils.Read(&Term, "pkg/data/Term")
	utils.Read(&VoteFor, "pkg/data/VoteFor")
	utils.Read(&Log, "pkg/data/Log")
	utils.Read(&CommitIndex, "pkg/data/CommitIndex")
	utils.Read(&LastApplied, "pkg/data/LastApplied")
	if Term > 0 {
		rf.setTerm(Term)
	}
	if VoteFor != "" {
		rf.setVoteFor(VoteFor)
	}
	if Log != nil {
		rf.log = Log
		rf.StartIndex = Log[0].Index
	}
	if CommitIndex > 0 {
		rf.commitIndex = CommitIndex
	}
	if LastApplied > 0 {
		rf.lastApplied = LastApplied
	}
	raw, err := ioutil.ReadFile("pkg/data/snap")
	if err != nil {
		log.Fatalln(err)
	}
	//就是从本地拿出快照再加上log，生成最新状态的状态机
	ns := &Snap{}
	json.Unmarshal(raw, ns)
	rf.Stroage = ns.Stroage
	for i := 0; i < len(rf.log); i++ {
		if rf.log[i].Index <= ns.Index {
			continue
		} else {
			rf.AppendToMachine(rf.log[i])
		}
	}

}

func (rf *Raft) AutoFlush() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		utils.Write(rf.currentTerm, "pkg/data/Term")
		utils.Write(rf.votedFor, "pkg/data/VoteFor")
		utils.Write(rf.log, "pkg/data/Log")
		utils.Write(rf.commitIndex, "pkg/data/CommitIndex")
		utils.Write(rf.lastApplied, "pkg/data/LastApplied")
	}
}
func (rf *Raft) GetLastLogIndexAndTerm() (int32, int32) {
	if len(rf.log) > 0 {
		return rf.log[len(rf.log)-1].Index, rf.log[len(rf.log)-1].Term
	}
	return -1, -1
}
func (rf *Raft) TocommitTomachine() {
	for {
	_:
		<-rf.SignToCommit
		//rf.lock.Lock()
		//defer rf.lock.Unlock()
		savedTerm := rf.currentTerm
		savedLastApplied := rf.lastApplied
		var entries []*models.LogEntry
		if rf.commitIndex > rf.lastApplied {
			entries = rf.log[rf.lastApplied-rf.StartIndex+1 : rf.commitIndex-rf.StartIndex+1]
			//rf.lastApplied = rf.commitIndex
		}
		log.Println("提交了数据到复制机上，数据为", entries, "提交开始的位置为", savedLastApplied, "CommitIndex为", rf.commitIndex, "当前任期为", savedTerm)

		for _, entry := range entries {
			rf.CommitChannel <- models.LogEntry{
				Command: entry.Command,
				Index:   entry.Index,
				Term:    savedTerm,
			}
		}
	}
}
func (rf *Raft) CommitToMachine() {
	for {
		log := <-rf.CommitChannel
		rf.lastApplied++
		if !rf.CommondIds[log.Command.Id] {
			rf.CommondIds[log.Command.Id] = true
			if log.Command.Type == "ADD" {
				rf.Stroage.PutKey(log.Command.Key, log.Command.Value, time.Millisecond*time.Duration(log.Command.TimeDuration))
			} else if log.Command.Type == "DELETE" {
				rf.Stroage.DeleteKey(log.Command.Key)
			}
		}
		if len(rf.log) >= 100000 {
			rf.lock.Lock()
			start := 0
			if rf.lastApplied-rf.StartIndex >= 500000 {
				start = 500000
			} else {
				start = int(rf.lastApplied - rf.StartIndex)
			}
			rf.StartIndex = rf.log[start].Index
			rf.Snap.Stroage = rf.Stroage
			rf.Snap.Index = rf.log[start].Index - 1
			rf.log = rf.log[start:]
			fmt.Println(rf.lastApplied)
			data, err := json.Marshal(rf.Snap)
			err = ioutil.WriteFile("pkg/data/snap", data, 0600)
			if err != nil {
				panic(err)
			}
			rf.lock.Unlock()
		}
	}
}

func (rf *Raft) AppendToMachine(log *models.LogEntry) {
	if !rf.CommondIds[log.Command.Id] {
		rf.CommondIds[log.Command.Id] = true
		if log.Command.Type == "ADD" {
			rf.Stroage.PutKey(log.Command.Key, log.Command.Value, time.Millisecond*time.Duration(log.Command.TimeDuration))
		} else if log.Command.Type == "DELETE" {
			rf.Stroage.DeleteKey(log.Command.Key)
		}
	}
}
func (rf *Raft) AppendCommond(command *models.Command) (*models.AppendCommandReply, error) {

	rf.lock.Lock()
	defer rf.lock.Unlock()
	//如果不是leader，返回false
	if rf.state != Leader {
		fmt.Println("我不是leader")
		return &models.AppendCommandReply{IsReceive: false}, nil
	}
	rf.log = append(rf.log, &models.LogEntry{Command: command, Term: rf.currentTerm, Index: rf.StartIndex + int32(len(rf.log))})
	log.Println("leader接受到了client的指令：", &command, "当前任期为", rf.currentTerm)
	rf.AppendToFollowerChannel <- struct{}{}
	return &models.AppendCommandReply{IsReceive: true}, nil

}

// 设置任期
func (rf *Raft) setTerm(term int32) {
	rf.currentTerm = term
}

// 设置为谁投票
func (rf *Raft) setVoteFor(id string) {
	rf.votedFor = id
}

// 设置当前领导者
func (rf *Raft) setCurrentLeader(leader string) {
	rf.currentLeader = leader
}

// 设置当前状态
func (rf *Raft) setStatus(state RaftState) {
	rf.state = state
}

// 投票累加
func (rf *Raft) voteAdd() {
	rf.vote++
}

// 设置投票数量
func (rf *Raft) setVote(num int) {
	rf.vote = num
}

// 心跳超时时间
func (rf *Raft) electionTimeout() time.Duration {
	if rf.state == New {
		rf.state = Follower
		return time.Duration(2500+rand.Intn(500)) * time.Millisecond
	}
	return time.Duration(300+rand.Intn(300)) * time.Millisecond
}

// 选举计时器
func (rf *Raft) runElectionTimer() {
	log.Println("开始选举计时")
	log.Println(rf.state, rf.currentTerm)
	rf.timeoutDuration = rf.electionTimeout()
	rf.lock.Lock()
	//termStarted := rf.currentTerm
	rf.lock.Unlock()

	ticker := time.NewTicker(50 * time.Millisecond)

	//在follower中，这通常在 rf 的生命周期内一直在后台运行
	defer ticker.Stop()
	for {
		<-ticker.C
		rf.lock.Lock()
		//如果我们没有收到领导人的消息或在超时期间没有投票给某人，则开始预投票，预投票成功后开始选举。
		if elapsed := time.Since(rf.electionResetEvent); elapsed >= rf.timeoutDuration && (rf.state == Follower || rf.state == Candidate) {
			log.Println("获取心跳超时，进行预投票操作")
			rf.becomeFollower()
			rf.PreRequestVote()
			//rf.lock.Unlock()
			//return
		}
		rf.lock.Unlock()
	}
}
func (rf *Raft) PreRequestVote() bool {
	savedCurrentTerm := rf.currentTerm
	IsOkToVote := false
	//同时向所有其他服务器发送 PreRequestVote RPC。
	votecount := 1
	for _, peerId := range rf.peerIds {
		if peerId == rf.id {
			continue
		}
		go func(peerId string) {
			rf.lock.Lock()
			lastLogIndex, lastLogTerm := rf.GetLastLogIndexAndTerm()
			rf.lock.Unlock()
			if rf.state != Follower {
				log.Println("该节点在预投票投票时状态更改为:", rf.state, "取消投票请求")
				return
			}
			args := models.RequestVoteArgs{
				Term:         rf.currentTerm,
				CandidateId:  rf.id,
				LastLogIndex: lastLogIndex,
				LastLogTerm:  lastLogTerm,
				Type:         "PreVote",
			}
			//GRPC通信
			client := rf.server.GetClient(peerId)
			reply, err := client.RequestVote(context.Background(), &args)
			if err != nil {
				//log.Println(err)
				return
			}
			rf.lock.Lock()
			defer rf.lock.Unlock()
			if reply.Term > savedCurrentTerm {
				rf.currentTerm = reply.Term
				rf.becomeFollower()
				log.Println("该节点在预投票时得到了大于自己任期的结果,预投票设为失败")
			} else if reply.Term == savedCurrentTerm && reply.VoteReply {
				votecount++
				if votecount >= (len(rf.peerIds)+1)/2 {
					log.Println("预投票成功，成为候选人，开始选举")
					rf.startElection()
				}
			}
			return
		}(peerId)
	}
	return IsOkToVote
}
func (rf *Raft) startElection() {
	log.Println("获取超时，开始选举")
	//更改状态为候选人，任期加1
	rf.setStatus(Candidate)
	rf.setTerm(rf.currentTerm + 1)
	savedCurrentTerm := rf.currentTerm
	//重设选举计时过期时间
	rf.electionResetEvent = time.Now()
	//为自己投票
	rf.setVoteFor(rf.id)
	rf.setVote(1)

	log.Println("成为候选人,当前任期为:", savedCurrentTerm)

	//同时向所有其他服务器发送 RequestVote RPC。
	for _, peerId := range rf.peerIds {
		if peerId == rf.id {
			continue
		}
		go func(peerId string) {
			rf.lock.Lock()
			lastLogIndex, lastLogTerm := rf.GetLastLogIndexAndTerm()
			rf.lock.Unlock()
			args := models.RequestVoteArgs{
				Term:         savedCurrentTerm,
				CandidateId:  rf.id,
				LastLogIndex: lastLogIndex,
				LastLogTerm:  lastLogTerm,
				Type:         "Vote",
			}
			//GRPC通信
			client := rf.server.GetClient(peerId)
			reply, err := client.RequestVote(context.Background(), &args)
			if err != nil {
				//log.Println(err)
				return
			}
			rf.lock.Lock()
			defer rf.lock.Unlock()
			if rf.state != Candidate {
				log.Println("该节点在请求投票时状态更改为:", rf.state, "取消投票请求")
				return
			}
			if reply.Term > savedCurrentTerm {
				rf.currentTerm = reply.Term

				rf.becomeFollower()

				log.Println("该节点在请求投票时得到了大于自己任期的结果,取消投票请求")
				return
			} else if reply.Term == savedCurrentTerm && reply.VoteReply {
				rf.voteAdd()
				if rf.vote >= (len(rf.peerIds)+1)/2 {
					rf.leaderstart()
				}
			}
			log.Println(reply)

		}(peerId)
	}

	rf.timeoutDuration = rf.electionTimeout()
}
func (rf *Raft) leaderstart() {
	log.Println("成为leader")

	rf.setStatus(Leader)
	for _, peerId := range rf.peerIds {
		if rf.id == peerId {
			continue
		}
		rf.nextIndex[peerId] = int32(len(rf.log)) + rf.StartIndex
		rf.matchIndex[peerId] = -1
	}
	//50ms发送心跳
	go func() {
		log.Println("发送心跳")
		timer := time.NewTimer(50 * time.Millisecond)
		defer timer.Stop()
		// Send periodic heartbeats, as long as still leader.
		for {
			tosend := false
			select {
			case <-timer.C:
				timer.Stop()
				timer.Reset(50 * time.Millisecond)
				tosend = true
			case _, ok := <-rf.AppendToFollowerChannel:
				if ok {
					tosend = true
				} else {
					return
				}
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(50 * time.Millisecond)

			}
			if tosend {
				rf.lock.Lock()
				if rf.state != Leader {
					rf.lock.Unlock()
					return
				}
				rf.leaderSendHeartbeats()
				rf.lock.Unlock()
			}
		}
	}()
}
func (rf *Raft) leaderSendHeartbeats() {
	savedCurrentTerm := rf.currentTerm
	for _, peerId := range rf.peerIds {
		if peerId == rf.id {
			continue
		}
		go func(peerId string) {
			client := rf.server.GetClient(peerId)
			rf.lock.Lock()
			nxt := rf.nextIndex[peerId]
			prevLogIndex := nxt - 1
			prevLogTerm := int32(-1)
			if prevLogIndex-rf.StartIndex >= 0 {
				prevLogTerm = rf.log[prevLogIndex-rf.StartIndex].Term
			}
			entries := rf.log[nxt-rf.StartIndex:]

			args := models.AppendEntriesArgs{
				Term:         savedCurrentTerm,
				LeaderId:     rf.id,
				PrevLogIndex: prevLogIndex,
				PrevLogTerm:  prevLogTerm,
				Entries:      entries,
				LeaderCommit: rf.commitIndex,
			}
			rf.lock.Unlock()
			reply, err := client.AppendEntries(context.Background(), &args)
			if err != nil {
				//log.Println(err)
				return
			}
			rf.lock.Lock()
			defer rf.lock.Unlock()
			if reply.Term > savedCurrentTerm {
				log.Println("发送心跳时，leader发现自己的任期小于回复节点的任期")
				rf.setTerm(reply.Term)

				rf.becomeFollower()

				return
			}
			if rf.state == Leader && savedCurrentTerm == reply.Term {
				if reply.Success {
					rf.nextIndex[peerId] = nxt + int32(len(entries))
					rf.matchIndex[peerId] = rf.nextIndex[peerId] - 1

					savedCommitIndex := rf.commitIndex
					for i := int(rf.commitIndex - rf.StartIndex + 1); i < len(rf.log); i++ {
						if rf.log[i].Term == rf.currentTerm {
							matchCount := 1
							for _, peerId := range rf.peerIds {
								if int(rf.matchIndex[peerId]) >= i+int(rf.StartIndex) {
									matchCount++
								}
							}
							if matchCount >= (len(rf.peerIds)+1)/2 {
								rf.commitIndex = int32(i) + rf.StartIndex
							}
						}
					}
					if rf.commitIndex != savedCommitIndex {
						log.Println("发送信号，可以向状态机提交索引")
						rf.SignToCommit <- struct{}{}
						//发送信号，可以向从机更新可提交索引
						rf.AppendToFollowerChannel <- struct{}{}
					}
				} else {
					if nxt == rf.StartIndex && !rf.Snapbool[peerId] {
						//在最前面都没找到，说明和主的状态差距过大，主把快照发过去
						rf.Snapbool[peerId] = true
						go rf.SendSnapToFollower(peerId)
					} else {
						//往前找，直到找到对应的
						rf.nextIndex[peerId] = nxt - 1
					}
				}
			}
		}(peerId)
	}
}

func (rf *Raft) SendSnapToFollower(peerId string) {
	for i := 0; i <= 3; i++ {
		client := rf.server.GetClient(peerId)
		data, _ := json.Marshal(rf.Snap)
		_, err := client.SendSnap(context.Background(), &models.SnapArgs{Data: data})
		if err != nil {
			continue
		} else {
			break
		}
	}
	rf.Snapbool[peerId] = false

}
func (rf *Raft) loadSnap(args *models.SnapArgs) *models.SnapReply {
	data := &Snap{}
	json.Unmarshal(args.Data, data)
	if data.Index < rf.Snap.Index || data.Index < rf.StartIndex+int32(len(rf.log)) {
		return &models.SnapReply{IsAccept: false}
	} else {
		rf.Snap = data
		rf.Stroage = data.Stroage
		rf.log = nil
		rf.StartIndex = data.Index + 1
		return &models.SnapReply{IsAccept: true}
	}

}

// AppendEntries
// 请求
// term	领导人的任期
// leaderId	领导人 ID 因此跟随者可以对客户端进行重定向（译者注：跟随者根据领导人 ID 把客户端的请求重定向到领导人，比如有时客户端把请求发给了跟随者而不是领导人）
// prevLogIndex	紧邻新日志条目之前的那个日志条目的索引
// prevLogTerm	紧邻新日志条目之前的那个日志条目的任期
// entries[]	需要被保存的日志条目（被当做心跳使用时，则日志条目内容为空；为了提高效率可能一次性发送多个）
// leaderCommit	领导人的已知已提交的最高的日志条目的索引
// 返回
// term	当前任期，对于领导人而言 它会更新自己的任期
// success	如果跟随者所含有的条目和 prevLogIndex 以及 prevLogTerm 匹配上了，则为 true
// AppendEntries 由领导人调用，用于日志条目的复制，同时也被当做心跳使用
func (rf *Raft) AppendEntries(args *models.AppendEntriesArgs) (*models.AppendEntriesReply, error) {
	//log.Println("接受心跳")
	rf.lock.Lock()
	defer rf.lock.Unlock()
	reply := &models.AppendEntriesReply{Term: rf.currentTerm, Success: false}
	if rf.state == Dead {
		return reply, nil
	}
	if args.Term > rf.currentTerm {
		log.Println("接受到心跳信号，信号的任期大于当前任期，当前节点变为追随节点")
		rf.setTerm(args.Term)

		rf.becomeFollower()

	}
	if args.Term == rf.currentTerm {
		if rf.state != Follower {
			rf.becomeFollower()
		}
		rf.electionResetEvent = time.Now()
		//找到第一个与leader位置索引不产生冲突的地方
		if args.PrevLogIndex == -1 || (args.PrevLogIndex < int32(len(rf.log))+rf.StartIndex && rf.log[args.PrevLogIndex-rf.StartIndex].Term == args.PrevLogTerm) {
			reply.Success = true
			rf.log = append(rf.log[:args.PrevLogIndex-rf.StartIndex+1], args.Entries...)
		}
		savedcommitIndex := rf.commitIndex
		if args.LeaderCommit > rf.commitIndex {
			if args.LeaderCommit > int32(len(rf.log)-1)+rf.StartIndex {
				rf.commitIndex = int32(len(rf.log)-1) + rf.StartIndex
			} else {
				rf.commitIndex = args.LeaderCommit
			}
			//提交索引位置改变，则发送提交信号
			if savedcommitIndex != rf.commitIndex {
				rf.SignToCommit <- struct{}{}
			}
		}
	}
	reply.Term = rf.currentTerm
	return reply, nil
}
func (rf *Raft) becomeFollower() {
	//log.Println("成为follower选举计时")
	//go rf.runElectionTimer()
	rf.state = Follower
	rf.votedFor = "-1"
	rf.setVote(0)
	//重设选举超时计时
	rf.electionResetEvent = time.Now()
	//重设超时时间
	rf.timeoutDuration = rf.electionTimeout()
}

// RequestVote RPC.
func (rf *Raft) RequestVote(args *models.RequestVoteArgs) (*models.RequestVoteReply, error) {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	reply := &models.RequestVoteReply{Term: rf.currentTerm, VoteReply: false}
	if rf.state == Dead {
		return reply, nil
	}
	//最近的日志索引和日志任期
	lastLogIndex, lastLogTerm := rf.GetLastLogIndexAndTerm()
	if args.Type == "PreVote" {
		if rf.currentTerm <= args.Term &&
			(rf.votedFor == "-1" || rf.votedFor == args.CandidateId) &&
			(args.LastLogTerm > lastLogTerm || (args.LastLogTerm == lastLogTerm && args.LastLogIndex >= lastLogIndex)) {
			reply.VoteReply = true
			log.Println("我给id为", args.CandidateId, "预投票", rf.state)
			rf.votedFor = args.CandidateId
			rf.electionResetEvent = time.Now()
		} else {
			//rf.electionResetEvent = time.Now()
			reply.VoteReply = false
		}
	} else if args.Type == "Vote" {
		//如果请求的任期小于等于当前节点任期，返回当前节点任期与false
		if args.Term <= rf.currentTerm {
			return reply, nil
		}
		//如果请求的任期大于当前节点任期，则当前节点任期设为请求任期，当前节点设为follower
		if args.Term > rf.currentTerm {
			log.Println("请求的任期大于当前节点任期，则当前节点任期设为请求任期，当前节点设为follower", args.Term, rf.currentTerm)
			rf.setTerm(args.Term)

			rf.becomeFollower()

		}
		log.Println(args.LastLogTerm, lastLogTerm, args.LastLogIndex, lastLogIndex, rf.votedFor)
		//如果当前未投票且请求的最新日志任期大于当前节点，或者在日志任期相等的情况下日志索引大于当前,则投票给请求节点
		if rf.currentTerm == args.Term &&
			(rf.votedFor == "-1" || rf.votedFor == args.CandidateId) &&
			(args.LastLogTerm > lastLogTerm || (args.LastLogTerm == lastLogTerm && args.LastLogIndex >= lastLogIndex)) {
			reply.VoteReply = true
			log.Println("我给id为", args.CandidateId, "投票", rf.state)
			rf.votedFor = args.CandidateId
			rf.electionResetEvent = time.Now()
		} else {
			//rf.electionResetEvent = time.Now()
			reply.VoteReply = false
		}
	}

	reply.Term = rf.currentTerm
	return reply, nil

}
