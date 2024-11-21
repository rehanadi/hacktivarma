package entity

import "time"

type Order struct {
	Id             string
	UserId         int
	DrugId         int
	Quantity       int
	Price          float64
	TotalPrice     float64
	PaymentMethod  string
	PaymentStatus  string
	PaymentAt      time.Time
	DeliveryStatus string
	DeliveryAt     time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
