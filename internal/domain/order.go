package domain

import "time"

type Order struct {
	ID          string
	CustomerID  string
	OrderDate   time.Time
	TotalAmount float64
	Products    []OrderProduct
}

type OrderProduct struct {
	ProductID string
	Quantity  int
	Price     float64
}