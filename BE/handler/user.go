package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)
 
func (h *Handler) CreateUser(c *gin.Context){
	var register user.Register
	errReq := c.ShouldBindJSON(&register)
	if (errReq != nil) {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	} 

	if (register.Password != register.ConfirmPassword){
		c.JSON(http.StatusBadRequest,"Password not match")
		return
	}

	hashedPassword,_ := helper.HashPassword(register.Password)
	user := user.User{
		Role : "student",
		Name : register.Name,
		Email: register.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : createdUser,
	})
}

func (h *Handler) GetCredentialUser(c *gin.Context){
	user,err := h.repo.GetUserByToken(c.Request.Header["Token"])
	if err != nil {
		c.JSON(http.StatusForbidden,"Forbidden")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : user,
	})
}

func (h *Handler) EditUser(c *gin.Context){
	var editUserRequest user.EditUserRequest
	errReq := c.ShouldBindJSON(&editUserRequest)
	if (errReq != nil) {
		c.JSON(http.StatusBadRequest,errReq)
		return
	} 

	if (editUserRequest.Password != editUserRequest.ConfirmPassword){
		c.JSON(http.StatusBadRequest,"Password not match")
		return
	}

	var user user.User
	user.Name = editUserRequest.Name
	user.Email = editUserRequest.Email
	user.Password,_ = helper.HashPassword(editUserRequest.Password)
	userUpdated,err := h.repo.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : userUpdated,
	})
}