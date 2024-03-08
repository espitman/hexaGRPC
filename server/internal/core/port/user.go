package port

import (
	"context"
	"github.com/espitman/hexaGRPC/internal/core/domain"
)

/**
 * UserService implemented by service.UserService interface
 */

type UserService interface {
	List(ctx *context.Context, skip, limit uint64) ([]domain.User, error)
}

/**
 * UserRepository implemented by repository.UserRepository interface
 */

type UserRepository interface {
	List(ctx *context.Context, skip, limit uint64) ([]domain.User, error)
}
