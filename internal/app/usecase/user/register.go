package user

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/lib/hash"
)

func (m *Module) Register(ctx context.Context, req entity.RegisterRequest) (err error) {
	//validate input
	if req.Username == "" || req.Password == "" || req.Email == "" || req.RoleID <= 0 {
		err = errs.BadRequest
		return
	}

	findUser, err := m.userRepo.FindUserByUsernameOrEmail(ctx, req.Username, req.Email)
	if err != nil {
		return
	}
	if findUser.Email != "" {
		err = errs.UsernameOrEmailAlreadyExist
		return
	}

	hashedPassword, err := hash.HashAndSalt(req.Password)
	if err != nil {
		return
	}
	req.Password = hashedPassword

	err = m.userRepo.InsertUser(ctx, req)

	return
}
