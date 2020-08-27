package user

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"

	"github.com/vincentwijaya/ventory-be/lib/jwt"
)

func (m *Module) Login(ctx context.Context, req entity.LoginRequest) (res entity.LoginResponse, err error) {
	//validate input
	if req.Email == "" || req.Password == "" {
		err = errs.BadRequest
		return
	}

	//TODO Find data to database then create token from that data
	claims := jwt.BuildTokenClaims(1, 540, "admin@admin.me", "admin")

	token, err := jwt.CreateToken(claims, m.jwtSecret)
	if err != nil {
		log.Errorf("failed to generate user token: %v", err)
		return
	}

	res.Token = token

	return
}
