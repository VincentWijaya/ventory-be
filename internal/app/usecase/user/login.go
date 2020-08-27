package user

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"

	"github.com/vincentwijaya/ventory-be/lib/hash"
	"github.com/vincentwijaya/ventory-be/lib/jwt"
)

func (m *Module) Login(ctx context.Context, req entity.LoginRequest) (res entity.LoginResponse, err error) {
	//validate input
	if req.Username == "" || req.Password == "" {
		err = errs.BadRequest
		return
	}

	findUser, err := m.userRepo.FindUserByUsernameOrEmail(ctx, req.Username, req.Username)
	if err != nil {
		return
	}
	if findUser.ID == 0 || findUser.Status != 1 {
		err = errs.Unauthorized
		return
	}

	comparePassword := hash.ComparePassword(findUser.Password, req.Password)
	if comparePassword == false {
		err = errs.Unauthorized
		return
	}

	claims := jwt.BuildTokenClaims(findUser.ID, 540, findUser.Email, findUser.Role)
	token, err := jwt.CreateToken(claims, m.jwtSecret)
	if err != nil {
		log.Errorf("failed to generate user token: %v", err)
		return
	}

	res.Token = token

	return
}
