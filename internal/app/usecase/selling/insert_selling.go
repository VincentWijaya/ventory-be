package selling

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) InsertSelling(ctx context.Context, req entity.InsertSellingRequest) (err error) {
	// validate request
	if req.ItemHistoryData.ItemID <= 0 || req.ItemHistoryData.ID <= 0 || req.ItemHistoryData.SellPrice <= 0 || req.ItemHistoryData.Stock <= 0 || req.QuantitySold <= 0 {
		err = errs.BadRequest
		return
	}

	findItem, err := m.itemRepo.FindItemByID(ctx, req.ItemHistoryData.ItemID)
	if err != nil {
		return err
	}
	if findItem.ItemName == "" || findItem.ID <= 0 {
		err = errs.NoData
		return
	}
	if findItem.Stock < req.QuantitySold {
		err = errs.InsufficientStock
		return
	}

	var (
		grossTotal float64
		netSell    float64
	)

	grossTotal = findItem.SellPrice * float64(req.QuantitySold)
	netSell = grossTotal - (findItem.BuyPrice * float64(req.QuantitySold))

	insertSelling := entity.Selling{
		ItemHistoryID: req.ItemHistoryData.ID,
		Quantity:      req.QuantitySold,
		GrossTotal:    grossTotal,
		NetSell:       netSell,
	}
	log.Infof("Insert selling: %+v", insertSelling)
	err = m.sellingRepo.InsertSelling(ctx, insertSelling)
	if err != nil {
		return
	}

	updateItem := entity.Item{
		ID:         findItem.ID,
		ItemName:   findItem.ItemName,
		BuyPrice:   findItem.BuyPrice,
		SellPrice:  findItem.SellPrice,
		Notes:      findItem.Notes,
		Stock:      findItem.Stock - req.QuantitySold,
		CategoryID: findItem.CategoryID,
	}
	log.Infof("Update item: %+v", updateItem)
	err = m.itemUsecase.UpdateItem(ctx, updateItem)
	if err != nil {
		return
	}

	return
}
