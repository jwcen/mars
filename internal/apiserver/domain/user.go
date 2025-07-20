package domain

import "time"

type User struct {
	Id        int64
	UserId    string
	Username  string
	Nickname  string
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
