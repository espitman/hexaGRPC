package grpc

import (
	greetpb "github.com/espitman/hexaGRPC/protos/protogen/greet"
	"google.golang.org/grpc"
)

func (s Server) registerServers(gs *grpc.Server) {
	greetpb.RegisterGreetServiceServer(gs, s.GreetHandler)
}
