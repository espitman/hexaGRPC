package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	GreetHandler *GreetHandler
	UserHandler  *UserHandler
}

func NewServer(greetHandler *GreetHandler, userHandler *UserHandler) *Server {
	return &Server{
		GreetHandler: greetHandler,
		UserHandler:  userHandler,
	}
}

func (s Server) Run() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting server:8080")

	gs := grpc.NewServer()
	s.registerServers(gs)

	if err := gs.Serve(listener); err != nil {
		panic(err)
	}

}
