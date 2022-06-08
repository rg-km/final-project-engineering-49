package repository

import (
	"belajar-golang/model/materi"
)

func (h *Repository) CountMateri() (int,error){
	// TO DO Safa Auliya
}

func (h *Repository) FindAllMateri() ([]materi.Materi, int ,error){
	// TO DO Safa Auliya
}

func (h *Repository) FindMateriByPage(Page int) ([]materi.Materi, int ,error) {
	// TO DO Safa Auliya
}

func (h *Repository) FindMateriByFilter(materiFilterRequest materi.MateriFilterRequest) ([]materi.Materi, int ,error) {
	// TO DO Safa Auliya
}

func (h *Repository) FindMateriByID(MateriID int) (materi.Materi, error) {
	// TO DO Safa Auliya
}

func (h *Repository) CreateMateri(materi materi.Materi) (materi.Materi, error) {
	// TO DO Safa Auliya
}

func (h *Repository) UpdateMateri(materi materi.Materi) (materi.Materi, error)  {
	// TO DO Safa Auliya
}

func (h *Repository) DeleteMateri(ID int) error {
	// TO DO Safa Auliya
}