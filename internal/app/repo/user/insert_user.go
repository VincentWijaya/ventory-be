package user

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (u *UserModule) InsertUser(ctx context.Context, data entity.RegisterRequest) error {
	q := u.MasterDB.Rebind(queries.InsertUser)
	_, err := u.MasterDB.Exec(ctx, q, data.Username, data.Email, data.Password, data.RoleID)
	return err
}
