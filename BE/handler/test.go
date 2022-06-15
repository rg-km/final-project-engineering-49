package handler

import (
	"belajar-golang/model/log"
	"belajar-golang/model/test"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTest(c *gin.Context){

	tests,err := h.repo.FindAllTest()
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : tests,
	})
}

func (h *Handler) GetTestByPage(c *gin.Context){

	pageN,errConv := strconv.Atoi(c.Param("page"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest,errConv.Error())
		return
	}

	tests,count,err := h.repo.FindTestByPage(pageN)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	data := test.TestResponse{
		Test : tests,
		Count: count,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : data,
	})
}
 
func (h *Handler) GetTestByMateri(c *gin.Context){

	materiN,errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest,errConv.Error())
		return
	}

	tests,err := h.repo.FindTestByMateri(materiN)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : tests,
	})
}

func (h *Handler) GetTestByID(c *gin.Context){

	TestID ,errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		c.JSON(http.StatusBadRequest,errConv.Error())
		return
	}

	test,err := h.repo.FindTestByID(TestID)
	if err != nil {
		c.JSON(http.StatusNotFound,err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : 200,
		"message" : "successfully",
		"data" : test,
	})
}

func (h *Handler) CreateTest(c *gin.Context) {

	var testRequest test.TestReqquest
	errReq := c.ShouldBind(&testRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	}

	test := test.Test{
		MateriID: testRequest.MateriID,
		Question: testRequest.Question,
		Answer1: testRequest.Answer1,
		Answer2: testRequest.Answer2,
		Answer3: testRequest.Answer3,
		Answer4: testRequest.Answer4,
		AnswerTrue: testRequest.AnswerTrue,
	}

	testCreated, err := h.repo.CreateTest(test)
	if err != nil {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : 200,
		"message" : "successfully",
		"data" : testCreated,
	})	
}

func (h *Handler) UpdateTest(c *gin.Context){
	//update is just delete and create again xixixixi :v
}

func (h *Handler) DeleteTest(c *gin.Context){
	ID,_ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteTest(ID)

	if (err != nil){
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
	})
}

func (h *Handler) SubmitTest(c *gin.Context){

	user,err := h.repo.GetUserByToken(c.Request.Header["Token"])
	if err != nil {
		c.JSON(http.StatusForbidden,"Forbidden")
		return
	}

	var logRequest log.LogRequest
	errReq := c.ShouldBindJSON(&logRequest)
	if errReq != nil {
		c.JSON(http.StatusBadRequest,errReq.Error())
		return
	}

	//for next version if lesson is available

	// isAccess := h.repo.IsAccessToLesson(user.ID,logRequest.MateriID)
	// if !isAccess {
	// 	c.JSON(http.StatusForbidden,"Forbidden")
	// 	return
	// }

	var score int = 0

	if logRequest.IsAnswer != 0 {
		score,err = h.repo.GetScore(logRequest.Answers,logRequest.TestID)
		if err!= nil {
			c.JSON(http.StatusBadRequest,err.Error())
			return
		}
	}

	log := log.Log{
		UserID: user.ID,
		MateriID: logRequest.MateriID,
		IsAnswer: logRequest.IsAnswer,
		Score: score,
	}

	err = h.repo.CreateOrUpdateLogForUser(log)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message" : "Successfully",
		"status" : 200,
		"data" : log,
	})
}

