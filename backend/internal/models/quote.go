package models

import "time"

type Quote struct {
	ID          int            `json:"quote_id"`
	QuoteCode   string         `json:"quote_code"`
	CustomerID  int            `json:"customer_id"`
	Status      string         `json:"status"`
	TotalAmount float64        `json:"total_amount"`
	Currency    string         `json:"currency"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Products    []QuoteProduct `json:"products"`
}

type QuoteProduct struct {
	ID           int       `json:"id"`
	QuoteID      int       `json:"quote_id"`
	ProductID    int       `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	PricePerUnit float64   `json:"price_per_unit"`
	TaxRate      float64   `json:"tax_rate"`
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
}
