package entity

import "time"

type User struct {
	Id        string
	Name      string
	Role      string
	Email     string
	Password  string
	Location  string
	CreatedAt time.Time
}

type UserStatistics struct {
	Total    int64
	Employee int64
	Customer int64
}
