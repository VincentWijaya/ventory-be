package item

import (
	"context"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

type (
	ItemRepository interface {
		InsertItem(ctx context.Context, item entity.Item) error
		UpdateItem(ctx context.Context, item entity.Item) error
		SoftDelete(ctx context.Context, itemID int64) error
		FindItem(ctx context.Context, lastID, dataPerPage int64) (res []entity.Item, err error)
		CountItem(ctx context.Context) (res entity.CountData, err error)
		InsertItemHistory(ctx context.Context, item entity.Item) error
	}
)

type Module struct {
	itemRepo ItemRepository
}

func New(itemRepo ItemRepository) *Module {
	return &Module{
		itemRepo: itemRepo,
	}
}
