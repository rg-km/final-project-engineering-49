package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCountOfTest(c *gin.Context) {
	// TO DO Nashihul Ibad
}

func (h *Handler) GetCountOfMateri(c *gin.Context) {
	// TO DO Nashihul Ibad
	count,err := h.repo.CountMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : count,
	})
}

func (h *Handler) GetCountOfStudent(c *gin.Context){
	//TO DO Nashihul Ibad
}
