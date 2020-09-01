package selling

import (
	"context"

	"github.com/vincentwijaya/ventory-be/constant/queries"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (s *SellingModule) InsertSelling(ctx context.Context, req entity.Selling) error {
	var err error
	q := s.MasterDB.Rebind(queries.InsertSelling)
	_, err = s.MasterDB.Exec(ctx, q, req.ItemHistoryID, req.GrossTotal, req.NetSell, req.Quantity)
	if err != nil {
		log.Errorf("InsertSelling: %+v", err)
	}

	return err
}
