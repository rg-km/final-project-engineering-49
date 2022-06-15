package repository

import (
	"belajar-golang/model/test"
	"fmt"
)

func (h *Repository) CountTest()(int,error){
	var count int64
	errCount := h.db.Model(&test.Test{}).Group("materi_id").Count(&count).Error
	if errCount != nil {
		return 0,errCount
	}
	return int(count),nil
}

func (h *Repository) FindAllTest()([]test.Test,error){
	var tests []test.Test
	errGet := h.db.Find(&tests).Error
	if errGet != nil {
		return tests,errGet
	}
	return tests,nil
}

func (h *Repository) FindTestByPage(Page int)([]test.Test,int,error){

	limit := 10
	offset := (Page - 1) * limit

	var tests []test.Test
	errGet := h.db.Limit(limit).Offset(offset).Find(&tests).Error
	if errGet != nil {
		return tests,0,errGet
	}

	var countRecord int64
	errCount := h.db.Model(&test.Test{}).Count(&countRecord).Error
	if errCount != nil {
		return tests,0,errCount
	}

	return tests,int(countRecord),nil
}

func (h *Repository) FindTestByMateri(MateriID int)([]test.Test,error){

	var tests []test.Test
	errGet := h.db.Where("materi_id = ?",MateriID).Find(&tests).Error
	if errGet != nil {
		return tests,errGet
	}

	return tests,nil
}

func (h *Repository) FindTestByID(TestID int)(test.Test,error) {
	var test test.Test
	err := h.db.Where("id = ?",TestID).First(&test).Error
	if err != nil {
		return test,err
	}
	return test,nil
}

func (h *Repository) CreateTest(test test.Test)(test.Test, error){
	err := h.db.Create(&test).Error
	if err != nil {
		return test,err
	}
	return test,nil
}

func (h *Repository) UpdateTest(test test.Test)(test.Test, error){
	err := h.db.Save(&test).Error
	if err != nil {
		return test,err
	}
	return test,nil
}

func (h *Repository) DeleteTest(TestID int)(error){
	
	var test test.Test
	err := h.db.Where("id = ?",TestID).First(&test).Error
	if err != nil {
		return err		
	}
	if test.ID == 0 {
		return fmt.Errorf("Data not found!")
	}
	err = h.db.Delete(&test).Error
	if err != nil {
		return err
	}
	return nil
}



