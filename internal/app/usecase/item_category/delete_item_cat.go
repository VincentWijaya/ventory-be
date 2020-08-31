package item_category

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
)

func (m *Module) DeleteItemCategory(ctx context.Context, categoryID int64) (err error) {
	// validate input
	if categoryID <= 0 {
		err = errs.BadRequest
		return
	}

	err = m.itemRepo.SoftDeleteItemCategory(ctx, categoryID)

	return err
}
