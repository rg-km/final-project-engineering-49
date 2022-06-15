package router

import (
	"belajar-golang/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setuprouter(Handler *handler.Handler) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Token", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	//nothing authorization
	r.POST("/register", Handler.CreateUser)
	r.POST("/login", Handler.Login)

	//user authorization
	r.Use(Handler.CheckUser)
	{
		r.GET("/user", Handler.GetCredentialUser)
		r.POST("/user/edit", Handler.EditUser)
		r.GET("/user/count")

		r.GET("/materis", Handler.GetAllMateri)
		r.GET("/materis/:page", Handler.GetMateriByPage)
		r.POST("/materis/filter", Handler.GetMateriByFilter)
		r.GET("/materi/:id", Handler.GetMateriByID)
		r.GET("/materi/count", Handler.GetCountOfMateri)

		r.GET("/test/materi/:id", Handler.GetTestByMateri)
		r.POST("/test/submit", Handler.SubmitTest)
		r.GET("/test/count", Handler.GetCountOfTest)
	}

	//admin authorization
	r.Use(Handler.CheckAdmin)
	{
		r.POST("/materi", Handler.CreateMateri)
		r.POST("/materi/update", Handler.UpdateMateri)
		r.GET("/materi/delete/:id", Handler.DeleteMateri)

		r.GET("/test/:id", Handler.GetTestByID)
		r.POST("/test", Handler.CreateTest)
		r.GET("/test/delete/:id", Handler.DeleteTest)
	}

	return r
}
