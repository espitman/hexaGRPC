package repository

import (
	"github.com/espitman/hexaGRPC/client/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserRepository struct {
}

/**
 * UserRepository implements port.UserRepository interface
 */

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (s *UserRepository) List(ctx *fiber.Ctx, skip, limit uint64) ([]domain.User, error) {
	var resp []domain.User
	resp = append(resp, domain.User{
		FName: "Saeed",
		LName: "Heidari",
	}, domain.User{
		FName: "John",
		LName: "Smith",
	})
	return resp, nil
}
