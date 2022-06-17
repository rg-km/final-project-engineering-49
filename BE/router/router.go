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
        AllowMethods:     []string{"POST","GET"},
        AllowHeaders:     []string{"Origin","Token","content-type"},
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
	
	r.Use(Handler.CheckUser)
	{
		r.GET("/user",  Handler.GetCredentialUser)
		r.GET("/materi",Handler.GetAllMateri)
		r.GET("/materi/id/:id",Handler.GetMateriByID)
		r.GET("/materi/:page/:limit",Handler.GetMateriByPage)
		r.GET("/materi/filter/:page/:limit",Handler.GetMateriByFilter)
		r.GET("/materi/count",Handler.GetCountMateri)
		r.GET("/test/materi/:id",Handler.GetTestByMateri)
		r.POST("/test/submit",Handler.SubmitTest)
		r.GET("/test/count",Handler.GetCountTest)
	}	

	r.Use(Handler.CheckAdmin)
	{
		r.POST("/materi",Handler.CreateMateri)
		r.POST("/materi/update",Handler.UpdateMateri)
		r.DELETE("/materi/delete/:id",Handler.DeleteMateri)
		r.POST("/test",Handler.CreateTest)
		r.DELETE("/test/delete/:id",Handler.DeleteTest)
	}

	return r
}