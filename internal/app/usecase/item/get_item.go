package item

import (
	"context"
	"math"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) GetItem(ctx context.Context, dataPerPage, lastID int64) (res entity.GetItemResponse, err error) {
	var totalPage int64
	totalPage = 1

	totalItem, err := m.itemRepo.CountItem(ctx)
	if err != nil {
		return
	}

	if totalItem.Total > dataPerPage {
		totalPage = int64(math.Ceil(float64(totalItem.Total) / float64(dataPerPage)))
	}
	if totalItem.Total == 0 {
		res.TotalPage = 0
		res.MaxDataPerPage = dataPerPage
		return
	}

	items, err := m.itemRepo.FindItem(ctx, lastID, dataPerPage)
	if err != nil {
		return
	}

	res.Items = items
	res.TotalPage = totalPage
	res.MaxDataPerPage = dataPerPage
	return
}
