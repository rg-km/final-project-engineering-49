package repository

import (
	"belajar-golang/model/test"
)

func (h *Repository) FindAllTest()([]test.Test,error){
	return nil,nil
}

func (h *Repository) FindTestByPage(Page int)([]test.Test,int,error){
	return nil,0,nil
}	

func (h *Repository) FindTestByID(TestID int)(test.Test,error) {
	return test.Test{},nil
}

func (h *Repository) CreateTest(test test.Test)(test.Test, error){
	return test,nil
}

func (h *Repository) UpdateTest(test test.Test)(test.Test, error){
	return test,nil
}

func (h *Repository) DeleteTest(TestID int)(error){
	return nil
}


