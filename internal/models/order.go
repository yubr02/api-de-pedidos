package models

import "time"

type Order struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customer_name"`
	Items        string    `json:"items"`
	Total        float64   `json:"total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
