package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	GreetHandler *GreetHandler
}

func NewServer(greetHandler *GreetHandler) *Server {
	return &Server{
		GreetHandler: greetHandler,
	}
}

func (s Server) Run() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	gs := grpc.NewServer()
	s.registerServers(gs)

	if err := gs.Serve(listener); err != nil {
		panic(err)
	}

	fmt.Println("starting server:8080")
}
