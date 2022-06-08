package enrollment

import "time"

type Enrollment struct {
	ID        int
	UserID    int
	LessonID  int
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}