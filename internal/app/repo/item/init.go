package item

import (
	"github.com/vincentwijaya/ventory-be/pkg/database"
)

type ItemModule struct {
	MasterDB database.DB
}

func New(masterDB database.DB) *ItemModule {
	return &ItemModule{
		MasterDB: masterDB,
	}
}
