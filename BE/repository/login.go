package repository

import (
	"belajar-golang/model/user"
	"fmt"
)

func (h *Repository) Login(login user.Login) (user.User, error){
	var user user.User
	err := h.db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return user,err
	}
	if (user.ID == 0){
		return user,fmt.Errorf("Not found")
	}
	return user,nil
}





