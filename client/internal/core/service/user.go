package service

import (
	"context"
	"github.com/espitman/hexaGRPC/client/internal/core/domain"
	"github.com/espitman/hexaGRPC/client/internal/core/port"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	UserClient port.UserGrpcClient
}

func NewUserService(userClient port.UserGrpcClient) *UserService {
	return &UserService{
		UserClient: userClient,
	}
}

func (s *UserService) List(ctx *context.Context, skip, limit uint64) ([]domain.User, error) {
	userResp := s.UserClient.GetUser(1)
	return []domain.User{domain.User{userResp.Fname, userResp.Lname}}, nil
}
