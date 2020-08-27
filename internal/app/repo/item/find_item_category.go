package item

import (
	"context"
	"database/sql"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) FindItemCategory(ctx context.Context, lastID, dataPerPage int64) (res []entity.ItemCategory, err error) {
	q := i.MasterDB.Rebind(queries.GetItemCategory)
	err = i.MasterDB.Get(ctx, &res, q, lastID, dataPerPage)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindItemCategory: %+v", err)
		return
	}

	return
}

func (i *ItemModule) CountItemCategory(ctx context.Context) (res entity.CountData, err error) {
	q := i.MasterDB.Rebind(queries.CountItemCategory)
	err = i.MasterDB.Select(ctx, &res, q)
	if err != nil {
		return
	}
	return
}
