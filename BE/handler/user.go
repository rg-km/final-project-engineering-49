package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAdmin(c *gin.Context){

	var register user.Register
	errReq := c.ShouldBindJSON(&register)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	if register.Password != register.ConfirmPassword {
		c.JSON(http.StatusBadRequest, "Password not match")
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
		c.JSON(http.StatusBadRequest, "email already exist")
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
		c.JSON(http.StatusBadRequest, err)
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
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	if register.Password != register.ConfirmPassword {
		c.JSON(http.StatusBadRequest, "Password not match")
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
		c.JSON(http.StatusBadRequest, "email already exist")
		return
	}

	hashedPassword, _ := helper.HashPassword(register.Password)
	user := user.User{
		Role:     "student",
		Name:     register.Name,
		Email:    register.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    createdUser,
	})

}

func (h *Handler) GetCredentialUser(c *gin.Context) {

	tknStr := c.Request.Header["Token"]

	user, err := h.repo.GetUserByToken(tknStr)
	if err != nil {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    user,
	})
}