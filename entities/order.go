package entity

import "time"

type Order struct {
	Id             string
	UserId         string
	DrugId         string
	Quantity       int
	Price          float64
	TotalPrice     float64
	PaymentMethod  string
	PaymentStatus  string
	PaymentAt      time.Time
	DeliveryStatus string
	DeliveredAt    time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserName       string
	DrugName       string
}
