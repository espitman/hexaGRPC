package grpc

import (
	greetpb "github.com/espitman/hexaGRPC/protos/protogen/greet"
	userpb "github.com/espitman/hexaGRPC/protos/protogen/user"
	"google.golang.org/grpc"
)

func (s Server) registerServers(gs *grpc.Server) {
	greetpb.RegisterGreetServiceServer(gs, s.GreetHandler)
	userpb.RegisterUserServiceServer(gs, s.UserHandler)
}
