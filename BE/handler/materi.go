package handler

import (
	"belajar-golang/model/materi"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func (h Handler) GetAllMateri(c *gin.Context) {

	materis, count, err := h.repo.FindAllMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======

func (h Handler) GetAllMateri(c *gin.Context){

	materis,count,err := h.repo.FindAllMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
		return
	}
 
	data := materi.MateriResponse{
		Materi: materis,
		Count: count,
	}

<<<<<<< HEAD
	data := materi.MateriResponse{
		Materi: materis,
		Count:  count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    data,
=======
	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : data,
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
	})
}

func (h *Handler) GetMateriByPage(c *gin.Context) {
	// TO DO Safa Auliya
	pageN, _ := strconv.Atoi(c.Param("page"))
	limitN, _ := strconv.Atoi(c.Param("limit"))
<<<<<<< HEAD

	materis, count, err := h.repo.FindMateriByPage(pageN, limitN)
=======
	
	materis, count, err := h.repo.FindMateriByPage(pageN,limitN)
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
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

	limit, _ := strconv.Atoi(c.Param("limit"))
	page, _ := strconv.Atoi(c.Param("page"))
	keyword := c.Query("keyword")

	var materiFilterRequest materi.MateriFilterRequest
	materiFilterRequest.Limit = limit
	materiFilterRequest.Page = page
	materiFilterRequest.Keyword = keyword

	materis, count, err := h.repo.FindMateriByFilter(materiFilterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
			"status":  400,
		})
		return
	}

	materi, err := h.repo.FindMateriByID(MateriID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
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
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, err.Error())
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
			"status":  400,
		})
		return
	}

	if materiRequest.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "course id null",
			"status":  400,
		})
		return
	}

	_, errFound := h.repo.FindMateriByID(materiRequest.ID)
	if errFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errFound.Error(),
			"status":  400,
		})
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
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
=======
		c.JSON(http.StatusBadRequest, err.Error())
>>>>>>> e8264e0d0d745420d12bec1f4858e2d7f8a313de
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
	})
}