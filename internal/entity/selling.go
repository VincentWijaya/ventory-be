package entity

import "time"

type Selling struct {
	ID            int64     `json:"sellingId" db:"id"`
	ItemHistoryID int64     `json:"itemHistoryId" db:"item_history_id"`
	Quantity      int64     `json:"quantity" db:"quantity"`
	GrossTotal    float64   `json:"grossTotal" db:"gross_total"`
	NetSell       float64   `json:"netSell" db:"net_sell"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
}
