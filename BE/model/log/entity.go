package log

import "time"

type Log struct {
	ID        int
	UserID    int
	MateriID  int
	IsAnswer  int
	Score 	  int
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}