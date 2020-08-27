package user

import (
	"context"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

type (
	UserRepository interface {
		FindUserByUsernameOrEmail(ctx context.Context, username, email string) (res entity.User, err error)
		FindAllUser(ctx context.Context) (res []entity.User, err error)
		InsertUser(ctx context.Context, data entity.RegisterRequest) error
	}
)

type Module struct {
	userRepo  UserRepository
	jwtSecret string
}

func New(user UserRepository, jwtSecret string) *Module {
	return &Module{
		userRepo:  user,
		jwtSecret: jwtSecret,
	}
}
