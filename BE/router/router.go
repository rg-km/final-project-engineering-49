package router

import (
	"belajar-golang/handler"

	"github.com/gin-gonic/gin"
)

func Setuprouter(Handler *handler.Handler) *gin.Engine {

	r := gin.Default()
	//nothing authorization
	r.POST("/register",Handler.CreateUser)
	r.POST("/login",Handler.Login)

	//user authorization
	r.Use(Handler.CheckUser)	

	r.GET("/user",  Handler.GetCredentialUser)
	r.POST("/user/edit",Handler.EditUser)

	r.GET("/materis",Handler.GetAllMateri)
	r.GET("/materis/:page",Handler.GetMateriByPage)
	r.POST("/materis/filter",Handler.GetMateriByFilter)
	r.GET("/materi/:id",Handler.GetMateriByID)
	r.POST("/materi/submit",Handler.SubmitMateri)
	r.GET("/materi/count",Handler.GetCountMateri)

	//admin authorization
	r.Use(Handler.CheckAdmin)

	r.POST("/materi",Handler.CreateMateri)
	r.POST("/materi/update",Handler.UpdateMateri)
	r.GET("/materi/delete/:id",Handler.DeleteMateri)
	return r
}