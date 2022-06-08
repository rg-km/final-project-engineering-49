package lesson

type LessonRequest struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	EnrollKey   string `binding:"required"`
}

type FilterLessonRequest struct {
	Name string `binding:"required"`
	Page int    `binding:"required"`
}
