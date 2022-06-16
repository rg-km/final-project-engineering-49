package log

type LogRequest struct {
	ID       int
	MateriID int `binding:"required"`
	Answers  []string
	TestID   []int
}