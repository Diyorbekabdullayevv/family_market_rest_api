package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Type        string         `json:"type" db:"type"`
	Description string         `json:"description" db:"description"`
	IsAvailable bool           `json:"is_available" db:"is_available"`

	CategoryID int            `json:"category_id" db:"category_id"`
	BrandID    int            `json:"brand_id" db:"brand_id"`

	UnitNumber       int            `json:"unit_number" db:"unit_number"`
	Barcode          string         `json:"barcode" db:"barcode"`
	ManufactureDate  time.Time      `json:"manufacture_date" db:"manufacture_date"`
	ReceiveDate      time.Time      `json:"receive_date" db:"receive_date"`
	ExpiryDate       time.Time      `json:"expiry_date" db:"expiry_date"`
	StockKeepingDate time.Time      `json:"stock_keeping_date" db:"stock_keeping_date"`

	Cost     float64       `json:"cost" db:"cost"`
	Discount float64       `json:"discount" db:"discount"`

	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Brand represents a product brand
type Brand struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}