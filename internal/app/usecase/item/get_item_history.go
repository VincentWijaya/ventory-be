package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) GetItemHistoryByItemID(ctx context.Context, itemID int64) (res []entity.ItemHistory, err error) {
	if itemID <= 0 {
		err = errs.BadRequest
		return
	}

	res, err = m.itemRepo.GetItemHistoryByItemID(ctx, itemID)
	if err != nil {
		return
	}

	resultLength := len(res)
	if resultLength > 0 {
		res = res[1:resultLength] // exclude current data
	} else {
		res = []entity.ItemHistory{}
	}

	return
}
