package user

import (
	"context"
	"database/sql"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (u *UserModule) FindUserByUsernameAndPassword(ctx context.Context, username, password string) (res entity.User, err error) {
	q := u.MasterDB.Rebind(queries.FindUserByUsernameAndPassword)
	err = u.MasterDB.Get(ctx, &res, q, username, password)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindUserByUsernameAndPassword: %+v", err)
		return
	}

	return
}

func (u *UserModule) FindAllUser(ctx context.Context) (res []entity.User, err error) {
	q := u.MasterDB.Rebind(queries.FindAllUser)
	err = u.MasterDB.Get(ctx, &res, q)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindAllUser: %+v", err)
		return
	}

	return
}
