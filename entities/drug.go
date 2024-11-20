package entity

import "time"

type Drug struct {
	Id          string
	Name        string
	Dose        float64
	Form        string
	Stock       int
	Price       float64
	Category    int
	ExpiredDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
