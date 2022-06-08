package materi

type MateriRequest struct {
	ID       int
	Title    string `binding:"required"`
	LessonID int
	Contain  string `binding:"required"`
	FileName string
	Status   int    `binding:"required"`
	Creator  string `binding:"required"`
	Image    string
}

type MateriFilterRequest struct {
	Keyword  string `binding:"required"`
	LessonID int
	Page     int `binding:"required"`
}
