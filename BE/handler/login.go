package handler

import (
	"belajar-golang/helper"
	"belajar-golang/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)
 
func (h *Handler) Login(c *gin.Context){

	var loginRequest user.Login
	errReq := c.ShouldBindJSON(&loginRequest)
	if (errReq != nil) {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	} 
 
	user, err := h.repo.Login(loginRequest)
	if err  != nil {
		c.JSON(http.StatusNotFound,err.Error())
		return 
	}

	match := helper.CheckPasswordHash(loginRequest.Password, user.Password)	
	if (!match){
		c.JSON(http.StatusBadRequest,"Password anda salah")
		return
	}

	user,err = h.repo.UpdateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : user,
	})
	return
}