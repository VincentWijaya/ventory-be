package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) InsertItem(ctx context.Context, item entity.Item) error {
	q := i.MasterDB.Rebind(queries.InsertItem)
	_, err := i.MasterDB.Exec(ctx, q, item.ItemName, item.CategoryID, item.BuyPrice,
		item.SellPrice, item.Stock, item.Notes)

	if err != nil {
		log.Errorf("InsertItem: %+v", err)
	}

	return err
}

func (i *ItemModule) UpdateItem(ctx context.Context, item entity.Item) error {
	q := i.MasterDB.Rebind(queries.UpdateItem)
	_, err := i.MasterDB.Exec(ctx, q, item.ItemName, item.CategoryID, item.BuyPrice,
		item.SellPrice, item.Stock, item.Notes, item.ID)

	if err != nil {
		log.Errorf("UpdateItem: %+v", err)
	}

	return err
}

func (i *ItemModule) SoftDelete(ctx context.Context, itemID int64) error {
	q := i.MasterDB.Rebind(queries.SoftDeleteItem)
	_, err := i.MasterDB.Exec(ctx, q, itemID)

	if err != nil {
		log.Errorf("SoftDelete: %+v", err)
	}

	return err
}
