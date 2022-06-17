package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
 
func (h *Handler) Login(c *gin.Context){

	var loginRequest user.Login
	errReq := c.ShouldBindJSON(&loginRequest)
	if (errReq != nil) {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	} 
 
	userFound, err := h.repo.Login(loginRequest)
	if err  != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return 
	}

	match := helper.CheckPasswordHash(loginRequest.Password, userFound.Password)	
	if (!match){
		c.JSON(http.StatusBadRequest,"Password anda salah")
		return
	}
	
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email: userFound.Email, 
		Role: userFound.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	var jwtKey = []byte("key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	}

	data := user.LoginResponse{
		Name: userFound.Name,
		Email: userFound.Email,
		Token : tokenString,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : data,
	})
	return	
}