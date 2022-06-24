package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCountStudent(c *gin.Context) {
	count, err := h.repo.GetCountStudent()
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
		"data":    count,
	})
}

func (h *Handler) GetCountMateri(c *gin.Context) {
	_, count, err := h.repo.FindAllMateri()
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
		"data":    count,
	})
}

func (h *Handler) GetCountTest(c *gin.Context) {
	count, err := h.repo.CountTestGroupByMateri()
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
		"data":    count,
	})
}
