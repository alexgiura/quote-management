package models

import "time"

type Product struct {
	ID           int       `json:"id"`
	ProductCode  string    `json:"product_code"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PricePerUnit float64   `json:"price_per_unit"`
	TaxRate      float64   `json:"tax_rate"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
