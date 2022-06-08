package test

import "time"

type Test struct {
	ID         int
	MateriID   int
	Question   string
	Answer1    string
	Answer2    string
	Answer3    string
	Answer4    string
	AnswerTrue string
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}