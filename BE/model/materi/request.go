package materi

type MateriRequest struct {
	ID       int
	Title    string `binding:"required"`
	Contain  string `binding:"required"`
	FileName string
	Creator  string `binding:"required"`
}

type MateriFilterRequest struct {
	Keyword  string `binding:"required"`
	LessonID int
	Page     int `binding:"required"`
}
