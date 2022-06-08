package repository

import (
	"belajar-golang/model/test"
)

func (h *Repository) FindAllTest()([]test.Test,error){

}

func (h *Repository) FindTestByPage(Page int)([]test.Test,int,error){

}

func (h *Repository) FindTestByID(TestID int)(test.Test,error) {
	
}

func (h *Repository) CreateTest(test test.Test)(test.Test, error){
	
}

func (h *Repository) UpdateTest(test test.Test)(test.Test, error){
	
}

func (h *Repository) DeleteTest(TestID int)(error){
	
}


