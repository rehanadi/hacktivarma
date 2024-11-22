package entity

import (
	"time"
)

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

type ReportOrder struct {
	Date               string
	DrugName           string
	TotalOrderAll      int
	TotalOrderPending  int
	TotalOrderSuccess  int
	TotalOrderFailed   int
	AmountOrderAll     float64
	AmountOrderPending float64
	AmountOrderSuccess float64
	AmountOrderFailed  float64
}
