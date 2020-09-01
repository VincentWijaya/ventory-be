package selling

import (
	"github.com/vincentwijaya/ventory-be/pkg/database"
)

type SellingModule struct {
	MasterDB database.DB
}

func New(masterDB database.DB) *SellingModule {
	return &SellingModule{
		MasterDB: masterDB,
	}
}
