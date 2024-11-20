package entity

import "time"

type User struct {
	Id        string
	Name      string
	Role      string
	Email     string
	Password  string
	CreatedAt time.Time
}
