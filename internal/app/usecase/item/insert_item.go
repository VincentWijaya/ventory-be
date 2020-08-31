package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) InsertItem(ctx context.Context, req entity.Item) (err error) {
	// validate input
	if req.ItemName == "" || req.CategoryID <= 0 || req.BuyPrice <= 0 || req.SellPrice <= 0 || req.Stock <= 0 {
		err = errs.BadRequest
		return
	}

	insertResult, err := m.itemRepo.InsertItem(ctx, req)
	if err != nil {
		return
	}

	insertedItemID, err := insertResult.LastInsertId()
	if err != nil {
		return err
	}

	req.ID = insertedItemID

	err = m.itemRepo.InsertItemHistory(ctx, req)

	return err
}
