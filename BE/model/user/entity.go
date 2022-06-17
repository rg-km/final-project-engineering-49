package user

import "time"

type User struct { 
	ID int
	Role string
	Name string
	Email string
	Password string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

