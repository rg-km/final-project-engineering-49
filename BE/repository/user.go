package repository

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (h *Repository) CreateAdmin()(error){
	val,_ := helper.HashPassword("admin123")
	user := user.User{
		Role : "admin",
		Name: "Admin for lms49",
		Email: "admin@gmail.com",
		Password: val,
	}
	err := h.db.Create(&user)
	if err != nil {
		return nil
	}
	return nil
}

func (h *Repository) CheckEmailIsExist(email string) bool{
	var user user.User
	h.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (h *Repository) GetUserByToken(token []string) (user.User, error){
	//get payload token
	claims := &Claims{}
	var jwtKey = []byte("key")
	jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	var user user.User
	h.db.Where("email = ?", claims.Email).First(&user)
	if user.ID == 0 {
		return user, fmt.Errorf("User not found")
	}
	return user, nil
}

func (h *Repository) CreateUser(user user.User) (user.User,error){
	err := h.db.Create(&user).Error
	if err != nil {
		return user,err
	}
	return user,nil
}

func (h *Repository) UpdateUser(user user.User) (user.User, error) {
	err := h.db.Save(&user).Error
	if err != nil {
		return user,err
	}
	return user,nil
}

func (h *Repository) IsLogin(token []string) (bool) {
	var user user.User
	h.db.Where("token = ?", token).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (h *Repository) IsAdmin(token []string) (bool) {
	var user user.User
	h.db.Where("token = ? and role = ?", token,"admin").First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

