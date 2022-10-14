package main

import (
	"raft/pkg"
)

var peers []string

func main() {
	peers = []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}
	pkg.Myserver = pkg.NewServer("127.0.0.1:8080", "127.0.0.1", "8080", peers)
	pkg.Myserver.Serve()
}
