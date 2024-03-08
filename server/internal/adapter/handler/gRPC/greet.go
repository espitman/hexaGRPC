package grpc

import (
	"context"
	"github.com/espitman/hexaGRPC/internal/core/port"
	greetpb "github.com/espitman/hexaGRPC/protos/protogen/greet"
	"log"
)

type GreetHandler struct {
	greetpb.UnimplementedGreetServiceServer
	userService port.UserService
}

func NewGreetHandler(userService port.UserService) *GreetHandler {
	return &GreetHandler{
		userService: userService,
	}
}

func (h GreetHandler) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	users, _ := h.userService.List(&ctx, 0, 1)

	log.Println("User name: ", req.UserName)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "us":
		greeting = "hello " + users[0].FName
	case "mx":
		greeting = "Hola " + users[0].FName
	default:
		greeting = "Hola/Hello " + users[0].FName
	}

	return &greetpb.GreetResponse{
		Result: greeting,
	}, nil
}
