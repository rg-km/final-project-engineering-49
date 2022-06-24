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

func (h *Handler) Login(c *gin.Context) {

	var loginRequest user.Login
	errReq := c.ShouldBindJSON(&loginRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
		return
<<<<<<< HEAD
	}

	userFound, err := h.repo.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
			"status":  400,
		})
=======
	} 
 
	userFound, err := h.repo.Login(loginRequest)
	if err  != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return 
	}

	match := helper.CheckPasswordHash(loginRequest.Password, userFound.Password)	
	if (!match){
		c.JSON(http.StatusBadRequest,"Password anda salah")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	match := helper.CheckPasswordHash(loginRequest.Password, userFound.Password)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password not match!",
			"status":  400,
		})
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
<<<<<<< HEAD
		Email: userFound.Email,
		Role:  userFound.Role,
=======
		Email: userFound.Email, 
		Role: userFound.Role,
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	var jwtKey = []byte("key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}

	data := user.LoginResponse{
<<<<<<< HEAD
		Name:  userFound.Name,
		Email: userFound.Email,
		Token: tokenString,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    data,
=======
		Name: userFound.Name,
		Email: userFound.Email,
		Token : tokenString,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : data,
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
	})
	return
}
