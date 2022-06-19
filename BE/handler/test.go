package handler

import (
	"belajar-golang/model/log"
	"belajar-golang/model/test"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTestByMateri(c *gin.Context) {
	// TO DO
	materiN, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest, errConv.Error())
		return
	}

	tests, err := h.repo.FindTestByMateri(materiN)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    tests,
	})
}

func (h *Handler) CreateTest(c *gin.Context) {
	// TO DO
	var testRequest test.TestReqquest
	errReq := c.ShouldBind(&testRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	test := test.Test{
		MateriID:   testRequest.MateriID,
		Question:   testRequest.Question,
		Answer1:    testRequest.Answer1,
		Answer2:    testRequest.Answer2,
		Answer3:    testRequest.Answer3,
		Answer4:    testRequest.Answer4,
		AnswerTrue: testRequest.AnswerTrue,
	}

	testCreated, err := h.repo.CreateTest(test)
	if err != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "successfully",
		"data":    testCreated,
	})
}

func (h *Handler) DeleteTest(c *gin.Context) {
	//TO DO
	ID, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteTest(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
	})
}

func (h *Handler) SubmitTest(c *gin.Context) {
	// TO DO
	user, err := h.repo.GetUserByToken(c.Request.Header["Token"])
	if err != nil {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	var logRequest log.LogRequest
	errReq := c.ShouldBindJSON(&logRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, errReq.Error())
		return
	}

	var score int = 0
	score, err = h.repo.GetScore(logRequest.Answers, logRequest.TestID)

	log := log.Log{
		UserID:   user.ID,
		MateriID: logRequest.MateriID,
		Score:    score,
	}

	logCreated, err := h.repo.SubmitTest(log)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"status":  200,
		"data":    logCreated,
	})
}
