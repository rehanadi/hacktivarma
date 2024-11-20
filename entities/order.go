package entity

import "time"

type Order struct {
	Id             string
	OrderNumber    string
	UserId         string
	TotalPrice     float64
	PaymentMethod  string
	PaymentStatus  string
	DeliveryStatus string
	DeliveryAt     time.Time
	PaymentAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrderDetail struct {
	Id        string
	OrderId   string
	DrugId    string
	Quantity  int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
