package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetCountMateri(c *gin.Context) {
	_,count,err := h.repo.FindAllMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : count,
	})
}

