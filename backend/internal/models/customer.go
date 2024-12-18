package models

import "time"

type Customer struct {
	ID           int       `json:"id"`
	CustomerCode string    `json:"customer_code"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
