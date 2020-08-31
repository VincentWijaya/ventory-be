package item_category

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) UpdateItemCategory(ctx context.Context, req entity.ItemCategory) (err error) {
	// validate input
	if req.CategoryName == "" || req.ID <= 0 {
		err = errs.BadRequest
		return
	}

	err = m.itemRepo.UpdateItemCategory(ctx, req)

	return err
}
