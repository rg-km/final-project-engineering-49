package handler

import (
	"belajar-golang/model/materi"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h Handler) GetAllMateri(c *gin.Context){

	materis,count,err := h.repo.FindAllMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}
 
	data := materi.MateriResponse{
		Materi: materis,
		Count: count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : data,
	})
}

func (h *Handler) GetMateriByPage(c *gin.Context) {
	// TO DO Safa Auliya
	pageN, errConv := strconv.Atoi(c.Param("page"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, errConv.Error())
		return
	}

	materis, count, err := h.repo.FindMateriByPage(pageN)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data := materi.MateriResponse{
		Materi: materis,
		Count:  count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    data,
	})
}

func (h *Handler) GetMateriByFilter(c *gin.Context) {
	// TO DO Safa Auliya
	var materiFilterRequest materi.MateriFilterRequest
	errReq := c.ShouldBindJSON(&materiFilterRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	materis, count, err := h.repo.FindMateriByFilter(materiFilterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data := materi.MateriResponse{
		Materi: materis,
		Count:  count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    data,
	})
}

func (h *Handler) GetMateriByID(c *gin.Context) {
	// TO DO Safa Auliya
	MateriID, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, errConv.Error())
		return
	}

	materi, err := h.repo.FindMateriByID(MateriID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "successfully",
		"data":    materi,
	})
}

func (h *Handler) CreateMateri(c *gin.Context) {
	// TO DO Safa Auliya
	var materiRequest materi.MateriRequest
	errReq := c.ShouldBindJSON(&materiRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	materi := materi.Materi{
		Title:    materiRequest.Title,
		Contain:  materiRequest.Contain,
		FileName: materiRequest.FileName,
		Creator:  materiRequest.Creator,
	}

	materiCreated, err := h.repo.CreateMateri(materi)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "successfully",
		"data":    materiCreated,
	})
}

func (h *Handler) UpdateMateri(c *gin.Context) {
	// TO DO Safa Auliya
	var materiRequest materi.MateriRequest
	errReq := c.ShouldBindJSON(&materiRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	if materiRequest.ID == 0 {
		c.JSON(http.StatusBadRequest, fmt.Errorf("materi id null"))
		return
	}

	_, errFound := h.repo.FindMateriByID(materiRequest.ID)
	if errFound != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("materi not found"))
		return
	}

	materi := materi.Materi{
		ID:       materiRequest.ID,
		Title:    materiRequest.Title,
		FileName: materiRequest.FileName,
		Contain:  materiRequest.Contain,
		Creator:  materiRequest.Creator,
	}

	materiUpdated, err := h.repo.UpdateMateri(materi)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "successfully",
		"data":    materiUpdated,
	})

}

func (h *Handler) DeleteMateri(c *gin.Context) {
	// TO DO Safa Auliya
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteMateri(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
	})
}