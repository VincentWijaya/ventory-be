package item

import (
	"context"
	"fmt"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

func (m *Module) GetItem(ctx context.Context, dataPerPage, lastID string) (res []entity.Item, err error) {
	totalItem, err := m.itemRepo.CountItem(ctx)
	if err != nil {
		return
	}

	fmt.Println("JUMLAH ITEM ===> ", totalItem.Total)
	return
}
