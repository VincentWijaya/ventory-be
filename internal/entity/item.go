package entity

import (
	"time"
)

type ItemCategory struct {
	ID           int64  `db:"id"`
	CategoryName string `db:"name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CountData struct {
	Total int64 `db:"total"`
}
