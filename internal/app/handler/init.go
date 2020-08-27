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
)

type Module struct {
	user       UsecaseUser
	middleware UsecaseMiddleware
}

func New(user UsecaseUser, middleware UsecaseMiddleware) *Module {
	return &Module{
		user:       user,
		middleware: middleware,
	}
}
