package service

import (
	"context"
	"github.com/espitman/hexaGRPC/internal/adapter/database/mongodb/repository"
	"github.com/espitman/hexaGRPC/internal/core/domain"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) List(ctx *context.Context, skip, limit uint64) ([]domain.User, error) {
	return s.UserRepository.List(ctx, skip, limit)
}
