package test

type TestReqquest struct {
	ID         int
	MateriID   int    `json:"materi_id" binding:"required"`
	Question   string `json:"question" binding:"required"`
	Answer1    string `json:"answer1" binding:"required"`
	Answer2    string `json:"answer2"  binding:"required"`
	Answer3    string `json:"answer3" binding:"required"`
	Answer4    string `json:"answer4" binding:"required"`
	AnswerTrue string `json:"answer_true" binding:"required"`
}