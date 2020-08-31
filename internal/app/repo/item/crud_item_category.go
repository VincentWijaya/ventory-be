package item

import (
	"context"
	"time"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (i *ItemModule) InsertItemCategory(ctx context.Context, data entity.ItemCategory) error {
	q := i.MasterDB.Rebind(queries.InsertItemCategory)
	_, err := i.MasterDB.Exec(ctx, q, data.CategoryName)

	if err != nil {
		log.Errorf("InsertItemCategory: %+v", err)
	}

	return err
}

func (i *ItemModule) UpdateItemCategory(ctx context.Context, data entity.ItemCategory) error {
	//Set date time to WIB
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	q := i.MasterDB.Rebind(queries.UpdateItemCategory)
	_, err = i.MasterDB.Exec(ctx, q, data.CategoryName, time.Now().In(loc), data.ID)

	if err != nil {
		log.Errorf("UpdateItemCategory: %+v", err)
	}

	return err
}

func (i *ItemModule) SoftDeleteItemCategory(ctx context.Context, itemCategoryID int64) error {
	//Set date time to WIB
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	q := i.MasterDB.Rebind(queries.SoftDeleteItemCategory)
	_, err = i.MasterDB.Exec(ctx, q, time.Now().In(loc), itemCategoryID)

	if err != nil {
		log.Errorf("SoftDeleteItemCategory: %+v", err)
	}

	return err
}
