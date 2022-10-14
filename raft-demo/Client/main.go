package main

import "Client/pkg"

func main() {
	peers := []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}
	pkg.NewServer("客户端", "", "", peers).Serve()
	select {}
}
