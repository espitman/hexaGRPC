package service

import (
	"context"
	"github.com/espitman/hexaGRPC/internal/core/domain"
	"github.com/espitman/hexaGRPC/internal/core/port"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) List(ctx *context.Context, skip, limit uint64) ([]domain.User, error) {
	return s.UserRepository.List(ctx, skip, limit)
}
