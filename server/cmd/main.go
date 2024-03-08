package main

import (
	"github.com/espitman/hexaGRPC/internal/adapter/database/mongodb/repository"
	grpc "github.com/espitman/hexaGRPC/internal/adapter/handler/gRPC"
	"github.com/espitman/hexaGRPC/internal/core/service"
)

func main() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	greetHandler := grpc.NewGreetHandler(userService)

	grpcServer := grpc.NewServer(greetHandler)
	grpcServer.Run()
}
