package lesson

import "time"

type Lesson struct {
	ID        int
	Name      string
	Description string
	EnrollKey string
	IsDelete  int
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}