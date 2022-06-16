package repository

import (
	"belajar-golang/model/log"
	"belajar-golang/model/materi"
	"belajar-golang/model/test"
)

func (h *Repository) FindTestByMateri(MateriID int) ([]materi.Materi, error) {
	//TO DO
	return nil, nil
}

func (h *Repository) CreateTest(test test.Test) (test.Test, error) {
	//TO DO
	return test, nil
}

func (h *Repository) DeleteTest(TestID int) error {
	//TO DO
	return nil
}

func (h *Repository) CountTestGroupByMateri() (int, error) {
	//TO DO
	return 0, nil
}

func (h *Repository) SubmitTest()(log.Log,error){
	//TO DO 
	return log.Log{},nil
}

func (h *Repository) GetScore(Answers []string,TestID []int)(int,error){
	//TO DO
	return 0,nil
}