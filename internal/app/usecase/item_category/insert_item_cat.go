package item_category

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) InsertItemCategory(ctx context.Context, req entity.ItemCategory) (err error) {
	// validate input
	if req.CategoryName == "" {
		err = errs.BadRequest
		return
	}

	err = m.itemRepo.InsertItemCategory(ctx, req)

	return err
}
