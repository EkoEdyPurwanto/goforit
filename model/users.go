package model

import "time"

type User struct {
	Id       string
	Username string
	Email    string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Users []User
