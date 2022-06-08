package test

type TestReqquest struct {
	ID         int
	MateriID   int    `binding:"required"`
	Question   string `binding:"required"`
	Answer1    string `binding:"required"`
	Answer2    string `binding:"required"`
	Answer3    string `binding:"required"`
	Answer4    string `binding:"required"`
	AnswerTrue string `binding:"required"`
}