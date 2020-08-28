package handler

import (
	"context"

	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/lib/jwt"
)

type (
	UsecaseUser interface {
		Login(ctx context.Context, request entity.LoginRequest) (entity.LoginResponse, error)
		Register(ctx context.Context, req entity.RegisterRequest) (err error)
	}
	UsecaseMiddleware interface {
		ValidateSession(ctx context.Context, sessionString string) (res *jwt.JWTValidateResponse, err error)
	}
	UsecaseItem interface {
		InsertItem(ctx context.Context, req entity.Item) (err error)
	}
)

type Module struct {
	user       UsecaseUser
	middleware UsecaseMiddleware
	item       UsecaseItem
}

func New(user UsecaseUser, middleware UsecaseMiddleware, item UsecaseItem) *Module {
	return &Module{
		user:       user,
		middleware: middleware,
		item:       item,
	}
}
