package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func (h *Handler) CreateAdmin(c *gin.Context) {
=======
func (h *Handler) CreateAdmin(c *gin.Context){
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de

	var register user.Register
	errReq := c.ShouldBindJSON(&register)
	if errReq != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, errReq.Error())
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	if register.Password != register.ConfirmPassword {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password not match",
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, "Password not match")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, "email already exist")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
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
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, err)
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
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
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, errReq.Error())
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	if register.Password != register.ConfirmPassword {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password not match",
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, "Password not match")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	existEmail := h.repo.CheckEmailIsExist(register.Email)
	if existEmail {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, "email already exist")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	hashedPassword, _ := helper.HashPassword(register.Password)
	user := user.User{
<<<<<<< HEAD
		Role:     "user",
=======
		Role:     "student",
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		Name:     register.Name,
		Email:    register.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, err)
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
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
<<<<<<< HEAD
	tknArr := strings.Split(tknStr[0], " ")
=======
	tknArr := strings.Split(tknStr[0]," ")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
	fixTkn := tknArr[1]

	user, err := h.repo.GetUserByToken(fixTkn)
	if err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusForbidden, "Forbidden")
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    user,
	})
<<<<<<< HEAD
}
=======
}
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
