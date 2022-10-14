package pkg

import (
	"Client/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var Myserver *Server

type Server struct {
	models.UnimplementedRaftRPCServer
	lock            sync.Mutex
	serverId        string
	peerIds         []string
	address         string
	port            string
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
	go func() {
		for {
			var Type, key, value string
			var Id int32
			fmt.Scan(&Type, &key, &value, &Id)
			fmt.Println(Type, key, value, Id)
			is := false
			for {
				if Type == "GET" {
					for _, peerId := range s.peerIds {
						if nowvalue, ok := s.GetKey(key, peerId); ok {
							is = true
							fmt.Println(nowvalue)
						}
					}
				} else {
					for _, peerId := range s.peerIds {
						if s.AppendCommand(Type, key, value, Id, peerId) {
							is = true
							fmt.Println("主节点已经收到了命令")
						}
					}
				}
				if is {
					break
				}
			}

		}
	}()
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
func (s *Server) GetKey(key, peerId string) (string, bool) {
	args := models.GetKeyArgs{Key: key}
	client := s.GetClient(peerId)
	reply, err := client.GetKey(context.Background(), &args)
	if err != nil {
		return "", false
	}
	return reply.Value, true
}
func (s *Server) AppendCommand(Type, key, value string, ID int32, peerId string) bool {
	Command := &models.Command{Type: Type, Key: key, Value: value, Id: ID}
	args := models.AppendCommandArgs{Command: Command}
	client := s.GetClient(peerId)
	reply, err := client.AppendCommand(context.Background(), &args)
	fmt.Println(reply)
	if err != nil {
		return false
	}
	if reply.IsReceive {
		return true
	}
	return false
}
