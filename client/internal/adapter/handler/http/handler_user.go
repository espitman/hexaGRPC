package http

import (
	"context"
	"github.com/espitman/hexaGRPC/client/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	ctx := context.Background()
	resp, _ := h.userService.List(&ctx, 0, 10)
	return c.JSON(resp)
}
