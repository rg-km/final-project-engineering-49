package repository

import (
	"belajar-golang/model/materi"
	"fmt"
)

func (h *Repository) CountMateri() (int, error) {
	// TO DO Safa Auliya
	return 0, nil
}

func (h *Repository) FindMateriByPage(Page int) ([]materi.Materi, int, error) {
	// TO DO Safa Auliya
	return nil, 0, nil
}

func (h *Repository) FindMateriByFilter(materiFilterRequest materi.MateriFilterRequest) ([]materi.Materi, int, error) {
	// TO DO Safa Auliya
	return nil, 0, nil
}

func (h *Repository) FindAllMateri() ([]materi.Materi, int, error) {
	// TO DO Safa Auliya
	var materis []materi.Materi
	errGet := h.db.Find(&materis).Error
	if errGet != nil {
		return materis, 0, errGet
	}

	var countRecord int64
	errCount := h.db.Model(&materi.Materi{}).Count(&countRecord).Error
	if errCount != nil {
		return materis, 0, errCount
	}
	return materis, int(countRecord), nil
}

func (h *Repository) FindMateriByID(MateriID int) (materi.Materi, error) {
	// TO DO Safa Auliya
	var materi materi.Materi

	err := h.db.Where("id = ?", MateriID).First(&materi).Error
	if err != nil {
		return materi, err
	}

	return materi, nil
}

func (h *Repository) CreateMateri(materi materi.Materi) (materi.Materi, error) {
	// TO DO Safa Auliya
	err := h.db.Create(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (h *Repository) UpdateMateri(materi materi.Materi) (materi.Materi, error) {
	// TO DO Safa Auliya
	err := h.db.Save(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (h *Repository) DeleteMateri(ID int) error {
	// TO DO Safa Auliya
	var materi materi.Materi
	err := h.db.Where("id = ?", ID).First(&materi).Error
	if err != nil {
		return err
	}
	if materi.ID == 0 {
		return fmt.Errorf("Data not found!")
	}
	err = h.db.Delete(&materi).Error
	if err != nil {
		return err
	}
	return nil
}
