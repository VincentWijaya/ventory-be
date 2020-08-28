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
		InsertItemCategory(ctx context.Context, data entity.ItemCategory) error
		UpdateItemCategory(ctx context.Context, data entity.ItemCategory) error
		SoftDeleteItemCategory(ctx context.Context, itemCategoryID int64) error
		FindItem(ctx context.Context, lastID, dataPerPage string) (res []entity.Item, err error)
		CountItem(ctx context.Context) (res entity.CountData, err error)
		FindItemByCategory(ctx context.Context, categoryID int64) (res []entity.Item, err error)
		FindItemByID(ctx context.Context, itemID int64) (res entity.Item, err error)
		FindItemCategory(ctx context.Context, lastID, dataPerPage int64) (res []entity.ItemCategory, err error)
		CountItemCategory(ctx context.Context) (res entity.CountData, err error)
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
