package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) InsertItemHistory(ctx context.Context, item entity.Item) error {
	q := i.MasterDB.Rebind(queries.InsertItemHistory)
	_, err := i.MasterDB.Exec(ctx, q, item.ID, item.BuyPrice,
		item.SellPrice, item.Stock, item.Notes)
	if err != nil {
		log.Errorf("InsertItemHistory: %+v", err)
	}

	return err
}

func (i *ItemModule) GetItemHistoryByItemID(ctx context.Context, itemID int64) (res []entity.ItemHistory, err error) {
	q := i.MasterDB.Rebind(queries.GetItemHistoryByItemID)
	err = i.MasterDB.Select(ctx, &res, q, itemID)
	if err != nil {
		log.Errorf("GetItemHistoryByItemID: %+v", err)
	}
	return
}
