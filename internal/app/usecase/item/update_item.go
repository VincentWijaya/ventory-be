package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) UpdateItem(ctx context.Context, req entity.Item) (err error) {
	// validate input
	if req.ItemName == "" || req.CategoryID <= 0 || req.BuyPrice <= 0 || req.SellPrice <= 0 || req.Stock <= 0 || req.ID <= 0 {
		err = errs.BadRequest
		return
	}

	err = m.itemRepo.InsertItemHistory(ctx, req)
	if err != nil {
		return
	}

	err = m.itemRepo.UpdateItem(ctx, req)
	return
}
