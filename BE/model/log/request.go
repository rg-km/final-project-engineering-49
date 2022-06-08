package log

type LogRequest struct {
	ID       int
	MateriID int `binding:"required"`
	IsAnswer int `binding:"required"`
	Answers  []string
	TestID   []int
}