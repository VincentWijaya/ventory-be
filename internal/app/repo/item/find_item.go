package item

import (
	"context"
	"database/sql"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) FindItem(ctx context.Context, lastID, dataPerPage int64) (res []entity.Item, err error) {
	q := i.MasterDB.Rebind(queries.GetItem)
	err = i.MasterDB.Select(ctx, &res, q, lastID, dataPerPage)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindItem: %+v", err)
		return
	}

	return
}

func (i *ItemModule) CountItem(ctx context.Context) (res entity.CountData, err error) {
	q := i.MasterDB.Rebind(queries.CountItem)
	err = i.MasterDB.Get(ctx, &res, q)
	if err != nil {
		log.Errorf("CountItem: %+v", err)
		return
	}
	return
}

func (i *ItemModule) FindItemByCategory(ctx context.Context, categoryID int64) (res []entity.Item, err error) {
	q := i.MasterDB.Rebind(queries.FindItemByCategoryID)
	err = i.MasterDB.Select(ctx, &res, q, categoryID)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindItemByCategory: %+v", err)
		return
	}

	return
}

func (i *ItemModule) FindItemByID(ctx context.Context, itemID int64) (res entity.Item, err error) {
	q := i.MasterDB.Rebind(queries.FindItemByID)
	err = i.MasterDB.Get(ctx, &res, q, itemID)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Errorf("FindItemByID: %+v", err)
		return
	}

	return
}
