package user

import (
	"github.com/vincentwijaya/ventory-be/pkg/database"
)

type UserModule struct {
	MasterDB database.DB
}

func New(masterDB database.DB) *UserModule {
	return &UserModule{
		MasterDB: masterDB,
	}
}
