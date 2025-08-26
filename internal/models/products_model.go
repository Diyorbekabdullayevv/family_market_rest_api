package models

import (
	"database/sql"
	"time"
)

type Products struct {
	Id          int        `json:"product_id" db:"product_id"`
	Name        string     `json:"product_name" db:"product_name"`
	Type        string     `json:"product_type" db:"product_type"`
	Description string     `json:"description" db:"description"`
	Category    string     `json:"category" db:"category"`
	Brand       string     `json:"brand" db:"brand"`
	IsAvailable bool       `json:"is_available" db:"is_available"`
	Codes       Codes      `json:"codes" db:"codes"`
	Dates       Dates      `json:"dates" db:"dates"`
	Pricing     Pricing    `json:"pricing" db:"pricing"`
	Timestamps  Timestamps `json:"timestamps" db:"timestamps"`
}

type Codes struct {
	UnitNumber int `json:"unit_number" db:"unit_number"`
	Barcode    int `json:"barcode" db:"barcode"`
}

type Dates struct {
	ManufactureDate  time.Time `json:"manufacture_date" db:"manufacture_date"`
	ReceiveDate      time.Time `json:"recieve_date" db:"recieve_date"`
	ExpiryDate       time.Time `json:"expiry_date" db:"expiry_date"`
	StockKeepingDate time.Time `json:"stock_keeping_date" db:"stock_keeping_date"`
}

type Pricing struct {
	Cost     float64 `json:"cost" db:"cost"`
	Discount float64 `json:"discount" db:"discount"`
}

type Timestamps struct {
	CreatedAt sql.NullString `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}
