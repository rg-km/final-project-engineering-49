package materi

import "time"

type Materi struct {
	ID        int
	Title     string
	Contain   string
	FileName  string
	Creator   string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}