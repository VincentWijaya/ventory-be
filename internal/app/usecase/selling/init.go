package selling

import (
	"context"

	"github.com/vincentwijaya/ventory-be/internal/entity"
)

type (
	SellingRepository interface {
		InsertSelling(ctx context.Context, req entity.Selling) error
	}
	ItemRepository interface {
		FindItemByID(ctx context.Context, itemID int64) (res entity.Item, err error)
	}
	ItemUsecase interface {
		UpdateItem(ctx context.Context, req entity.Item) (err error)
	}
)

type Module struct {
	sellingRepo SellingRepository
	itemRepo    ItemRepository
	itemUsecase ItemUsecase
}

func New(sellingRepo SellingRepository, itemRepo ItemRepository, itemUsecase ItemUsecase) *Module {
	return &Module{
		sellingRepo: sellingRepo,
		itemRepo:    itemRepo,
		itemUsecase: itemUsecase,
	}
}
