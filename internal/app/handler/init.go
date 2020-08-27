package handler

import (
	"context"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

type (
	UsecaseUser interface {
		Login(ctx context.Context, request entity.LoginRequest) (entity.LoginResponse, error)
	}
)

type Module struct {
	user UsecaseUser
}

func New(user UsecaseUser) *Module {
	return &Module{
		user: user,
	}
}
