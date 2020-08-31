package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) InsertItemHistory(ctx context.Context, item entity.Item) error {
	q := i.MasterDB.Rebind(queries.InsertItemHistory)
	_, err := i.MasterDB.Exec(ctx, q, item.ItemName, item.CategoryID, item.BuyPrice,
		item.SellPrice, item.Stock, item.Notes, item.UpdatedAt)
	if err != nil {
		log.Errorf("InsertItemHistory: %+v", err)
	}

	return err
}
