package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
)

func (m *Module) DeleteItem(ctx context.Context, itemID int64) error {
	if itemID <= 0 {
		err := errs.BadRequest
		return err
	}

	err := m.itemRepo.SoftDelete(ctx, itemID)
	return err
}
