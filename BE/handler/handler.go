package handler

import (
	"belajar-golang/repository"
)

type Handler struct {
	repo *repository.Repository
}

func NewRepo(repo *repository.Repository) *Handler{
	return &Handler{repo}
}

