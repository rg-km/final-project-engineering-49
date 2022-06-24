package handler

import (
	"belajar-golang/model/materi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllMateri(c *gin.Context) {

	materis, count, err := h.repo.FindAllMateri()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}
 
	data := materi.MateriResponse{
		Materi: materis,
		Count: count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    data,
	})
}

func (h *Handler) GetMateriByPage(c *gin.Context) {
	// TO DO Safa Auliya
	pageN, _ := strconv.Atoi(c.Param("page"))
	limitN, _ := strconv.Atoi(c.Param("limit"))

	materis, count, err := h.repo.FindMateriByPage(pageN, limitN)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  400,
		})
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