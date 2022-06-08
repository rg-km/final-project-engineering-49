package repository

import (
	"gorm.io/gorm"
)
 
type Repository struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *Repository {
	return &Repository{db}
}


