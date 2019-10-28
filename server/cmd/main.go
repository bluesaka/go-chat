package main

import "chat/server"

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")
	s.Start()
}
