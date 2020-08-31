package item_category

import (
	"context"
	"math"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) GetItemCategory(ctx context.Context, dataPerPage, lastID int64) (res entity.GetItemCategoryResponse, err error) {
	var totalPage int64
	totalPage = 1

	countCategory, err := m.itemRepo.CountItemCategory(ctx)
	if err != nil {
		return
	}

	if countCategory.Total > dataPerPage {
		totalPage = int64(math.Ceil(float64(countCategory.Total) / float64(dataPerPage)))
	}
	if countCategory.Total == 0 {
		res.Category = []entity.ItemCategory{}
		res.TotalPage = 0
		res.MaxDataPerPage = dataPerPage
		return
	}

	category, err := m.itemRepo.FindItemCategory(ctx, lastID, dataPerPage)
	if err != nil {
		return
	}

	res.Category = category
	res.TotalPage = totalPage
	res.MaxDataPerPage = dataPerPage
	return
}
