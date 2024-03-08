package grpcClient

import (
	"context"
	userpb "github.com/espitman/hexaGRPC/protos/protogen/user"
	"google.golang.org/grpc"
	"log"
)

type UserClient struct {
	UserServiceClient userpb.UserServiceClient
}

/**
 * UserClient implements port.UserGrpcClient interface
 */

func NewUserClient() *UserClient {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &UserClient{
		UserServiceClient: userpb.NewUserServiceClient(cc),
	}
}

func (uc *UserClient) GetUser(id uint32) *userpb.GetResponse {
	res, err := uc.UserServiceClient.Get(context.Background(), &userpb.GetRequest{
		Id: id,
	})
	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}
	return res
}
