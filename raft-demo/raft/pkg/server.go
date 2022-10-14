package pkg

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"raft/models"

	"sync"
	"syscall"
	"time"
)

var Myserver *Server

type Server struct {
	models.UnimplementedRaftRPCServer
	lock            sync.Mutex
	serverId        string
	peerIds         []string
	address         string
	port            string
	rf              *Raft
	peerClientConns map[string]*grpc.ClientConn
}

func NewServer(serverId, address, port string, peerIds []string) *Server {
	s := new(Server)
	s.serverId = serverId
	s.peerIds = peerIds
	s.address = address
	s.port = port
	s.peerClientConns = make(map[string]*grpc.ClientConn)
	return s
}

func (s *Server) Serve() {
	var start chan struct{}
	start = make(chan struct{})
	s.rf = NewRaft(s.serverId, s.address, s.port, s.peerIds, start)
	s.rf.server = s
	server, err := s.Start(start)
	if err != nil {
		fmt.Println("service started listen on", s.address+":"+s.port)
	}
	if err != nil {
		panic(fmt.Sprintf("start server failed : %v", err))
	}
	go func() {
		for {
			var tmp string
			fmt.Scanf("%s", &tmp)
			if tmp == "A" {
				s.rf.AppendCommond(&models.Command{
					Type:  "ADD",
					Key:   "你好",
					Value: "你吃饭了吗",
				})
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sign := <-c
		switch sign {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			server.Stop()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
func (s *Server) GetClient(ip string) models.RaftRPCClient {
	s.lock.Lock()
	defer s.lock.Unlock()
	if v, ok := s.peerClientConns[ip]; ok {
		return models.NewRaftRPCClient(v)
	} else {
		conn, err := grpc.Dial(ip, grpc.WithInsecure())
		if err != nil {
			log.Println(err)
		}
		s.peerClientConns[ip] = conn
		return models.NewRaftRPCClient(conn)
	}
}
func (s *Server) Start(startchan chan struct{}) (*grpc.Server, error) {
	server := grpc.NewServer()

	models.RegisterRaftRPCServer(server, new(Server))
	lis, err := net.Listen("tcp", s.address+":"+s.port)
	if err != nil {
		return nil, err
	}
	startchan <- struct{}{}
	go func() {
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return server, nil
}
func (s *Server) RequestVote(ctx context.Context, req *models.RequestVoteArgs) (res *models.RequestVoteReply, err error) {
	return Myserver.rf.RequestVote(req)
}
func (s *Server) AppendEntries(ctx context.Context, req *models.AppendEntriesArgs) (res *models.AppendEntriesReply, err error) {
	return Myserver.rf.AppendEntries(req)
}
func (s *Server) AppendCommand(ctx context.Context, req *models.AppendCommandArgs) (res *models.AppendCommandReply, err error) {
	return Myserver.rf.AppendCommond(req.Command)
}
func (s *Server) GetKey(ctx context.Context, req *models.GetKeyArgs) (res *models.GetKeyReply, err error) {
	return &models.GetKeyReply{IsOk: true, Value: Myserver.rf.Stroage.GetKey(req.Key)}, nil
}
func (s *Server) SendSnap(ctx context.Context, req *models.SnapArgs) (res *models.SnapReply, err error) {
	return Myserver.rf.loadSnap(req), nil
}
