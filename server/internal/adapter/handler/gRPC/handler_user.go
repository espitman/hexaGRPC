package grpc

import (
	"context"
	"github.com/espitman/hexaGRPC/internal/core/port"
	userpb "github.com/espitman/hexaGRPC/protos/protogen/user"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) Get(ctx context.Context, req *userpb.GetRequest) (*userpb.GetResponse, error) {

	userId := req.Id

	users, _ := h.userService.List(&ctx, 0, 1)

	return &userpb.GetResponse{
		Fname: users[userId].FName,
		Lname: users[userId].LName,
	}, nil
}
