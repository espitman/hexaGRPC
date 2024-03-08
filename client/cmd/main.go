package main

import (
	"github.com/espitman/hexaGRPC/client/internal/adapter/grpcClient"
	"github.com/espitman/hexaGRPC/client/internal/adapter/handler/http"
	"github.com/espitman/hexaGRPC/client/internal/core/service"
)

func main() {

	userClient := grpcClient.NewUserClient()
	userService := service.NewUserService(userClient)
	userHandler := http.NewUserHandler(userService)

	httpServer := http.NewServer("3000", userHandler)
	httpServer.Run()
}
