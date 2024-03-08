package port

import (
	"context"
	"github.com/espitman/hexaGRPC/client/internal/core/domain"
	userpb "github.com/espitman/hexaGRPC/protos/protogen/user"
)

/**
 * UserService implemented by service.UserService
 */

type UserService interface {
	List(ctx *context.Context, skip, limit uint64) ([]domain.User, error)
}

/**
 * UserRepository implemented by repository.UserRepository
 */

//type UserRepository interface {
//	List(ctx *context.Context, skip, limit uint64) ([]domain.User, error)
//}

/**
 * UserGrpcClient implemented by grpcClient.UserClient
 */

type UserGrpcClient interface {
	GetUser(id uint32) *userpb.GetResponse
}
