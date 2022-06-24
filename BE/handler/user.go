package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAdmin(c *gin.Context) {

	var register user.Register
	errReq := c.ShouldBindJSON(&register)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
		return
	}

	if register.Password != register.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password not match",
			"status":  400,
		})
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
			"status":  400,
		})
		return
	}

	hashedPassword, _ := helper.HashPassword(register.Password)
	user := user.User{
		Role:     "admin",
		Name:     register.Name,
		Email:    register.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    createdUser,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {

	var register user.Register
	errReq := c.ShouldBindJSON(&register)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
		return
	}

	if register.Password != register.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password not match",
			"status":  400,
		})
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
			"status":  400,
		})
		return
	}

	hashedPassword, _ := helper.HashPassword(register.Password)
	user := user.User{
		Role:     "user",
		Name:     register.Name,
		Email:    register.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    createdUser,
	})

}

func (h *Handler) GetCredentialUser(c *gin.Context) {

	tknStr := c.Request.Header["Authorization"]
	tknArr := strings.Split(tknStr[0], " ")
	fixTkn := tknArr[1]

	user, err := h.repo.GetUserByToken(fixTkn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    user,
	})
}
