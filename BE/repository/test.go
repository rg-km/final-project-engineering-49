package repository

import (
	"belajar-golang/model/log"
	"belajar-golang/model/test"
	"fmt"
	"time"
)

func (h *Repository) FindTestByMateri(MateriID int) ([]test.Test, error) {

	var sqlStatement string
	var tests []test.Test
	sqlStatement = `SELECT id,materi_id,question,answer1,answer2,answer3,answer4,answer_true
	FROM tests WHERE materi_id = ?`
	rows, err:= h.db.Query(sqlStatement,MateriID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var testTemp test.Test
		err := rows.Scan(&testTemp.ID,&testTemp.MateriID,&testTemp.Question,&testTemp.Answer1,&testTemp.Answer2,&testTemp.Answer3,&testTemp.Answer4,&testTemp.AnswerTrue)
		if err != nil {
			return nil,err
		}
		newTest := test.Test{
			ID : testTemp.ID,
			MateriID: testTemp.MateriID,
			Question: testTemp.Question,
			Answer1: testTemp.Answer1,
			Answer2: testTemp.Answer2,
			Answer3: testTemp.Answer3,
			Answer4: testTemp.Answer4,
			AnswerTrue: testTemp.AnswerTrue,
		}
		tests = append(tests, newTest)
	}
	rows.Close()
	return tests,nil

}

func (h *Repository) CreateTest(test test.Test) (test.Test, error) {
	//TO DO
	sqlStmt := `INSERT INTO tests (materi_id, question, answer1, answer2, answer3, answer4, answer_true,created_at, updated_at) 
	VALUES (?,?,?,?,?,?,?,?,?);`
	datetime := time.Now()  
	_, err := h.db.Exec(sqlStmt, test.MateriID, test.Question ,test.Answer1, test.Answer2,test.Answer3,test.Answer4,test.AnswerTrue,datetime,datetime);
	if err != nil {
		return test, err
	}
	return test,nil
}

func (h *Repository) DeleteTest(TestID int) error {
	sqlStmt := `DELETE FROM tests WHERE id = ?`
	_, err := h.db.Exec(sqlStmt, TestID)
	if err != nil {
		return err
	}
	return nil
}

func (h *Repository) CountTestGroupByMateri() (int, error) {
	var count int
	rows, err := h.db.Query("SELECT materi_id FROM tests GROUP BY materi_id")
	for rows.Next() {
		count++
	}
	if err != nil {
		return 0,err
	}
	return count, nil
}

func (h *Repository) SubmitTest(log log.Log)(log.Log,error){
	sqlStmt := `INSERT INTO logs (user_id, materi_id, score, created_at, updated_at) 
	VALUES (?,?,?,?,?);`
	datetime := time.Now()  
	_, err := h.db.Exec(sqlStmt, log.UserID,log.MateriID,log.Score,datetime,datetime);
	if err != nil {
		return log, err
	}
	return log,nil
}

func (h *Repository) GetScore(Answers []string,TestID []int)(int,error){
	if len(Answers) != len(TestID) {
		return 0, fmt.Errorf("length answer and test not match")
	}
	var answerTrue int
	var score int
	for i := 0; i<len(TestID); i++ {
		var testAnswerTrue string

		sqlStatement := `SELECT answer_true
		FROM tests WHERE id = ?`

		row := h.db.QueryRow(sqlStatement, TestID[i])
		err := row.Scan(&testAnswerTrue)
		if err != nil {
			continue
		}
		if testAnswerTrue != Answers[i] {
			continue
		}
		answerTrue++
		score = answerTrue
	}
	score = score*100/len(TestID)
	return score,nil
}