package log

type LogRequest struct {
	MateriID int      `json:"materi_id" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
	TestID   []int    `json:"test_id" binding:"required"`
}