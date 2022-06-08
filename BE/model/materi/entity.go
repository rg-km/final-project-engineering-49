package materi

import "time"

type Materi struct {
	ID        int
	LessonID  int
	Title     string
	Contain   string
	FileName  string
	Status    int
	Creator   string
	Image	  string
	IsDelete  int
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}