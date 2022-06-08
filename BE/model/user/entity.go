package user

import "time"

type User struct { 
	ID int
	Role string
	Name string
	Email string
	Password string
	Token string 
	CreatedAt time.Time
	UpdatedAt time.Time
}

