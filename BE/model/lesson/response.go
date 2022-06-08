package lesson

type LessonResponse struct {
	Lesson []Lesson
	Count  int
}

type LessonByUser struct {
	LessonID    int
	Name        string
	Description string
	UserID      int
}

type LessonByUserResponse struct {
	Lesson []LessonByUser
	Count  int
}