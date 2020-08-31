package entity

import (
	"time"
)

type ItemCategory struct {
	ID           int64     `json:"id" db:"id"`
	CategoryName string    `json:"categoryName" db:"category_name"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

type CountData struct {
	Total int64 `db:"total"`
}

type Item struct {
	ID           int64     `json:"id" db:"id"`
	ItemName     string    `json:"itemName" db:"item_name"`
	CategoryID   int64     `json:"categoryId" db:"category_id"`
	BuyPrice     float64   `json:"buyPrice" db:"buy_price"`
	SellPrice    float64   `json:"sellPrice" db:"sell_price"`
	Stock        int64     `json:"stock" db:"stock"`
	Notes        string    `json:"notes" db:"notes"`
	CategoryName string    `json:"categoryName" db:"category_name"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

type ItemHistory struct {
	ItemID       int64     `db:"item_id"`
	BuyPrice     float64   `json:"buyPrice" db:"buy_price"`
	SellPrice    float64   `json:"sellPrice" db:"sell_price"`
	Stock        int64     `json:"stock" db:"stock"`
	Notes        string    `json:"notes" db:"notes"`
	CategoryName string    `json:"categoryName" db:"category_name"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
