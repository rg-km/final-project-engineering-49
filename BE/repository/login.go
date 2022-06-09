package repository

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"fmt"
	"strconv"
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


func (h *Repository) UpdateToken(user user.User) (user.User, error){

	token := helper.String(32)+strconv.Itoa(user.ID);
	user.Token = token
	h.db.Save(&user)

	return user,nil
}



