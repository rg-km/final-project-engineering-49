package repository

import (
	"belajar-golang/model/user"
)

func (h *Repository) CreateAdmin()(error){
	// TO DO Nashihul Ibad
}

func (h *Repository) GetUserByToken(token []string) (user.User, error){
	// TO DO Nashihul Ibad
}

func (h *Repository) CreateUser(user user.User) (user.User,error){
	// TO DO Nashihul Ibad
}

func (h *Repository) UpdateUser(user user.User) (user.User, error) {
	// TO DO Nashihul Ibad
}

func (h *Repository) IsLogin(token []string) (bool) {
	// TO DO Nashihul Ibad
}

func (h *Repository) IsAdmin(token []string) (bool) {
	// TO DO Nashihul Ibad
}

