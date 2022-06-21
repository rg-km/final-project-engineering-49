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
	// r.POST("/register", Handler.CreateAdmin)
	r.POST("/login", Handler.Login)
	
	r.Use(Handler.CheckUser)
	{
		r.GET("/user",  Handler.GetCredentialUser)
		r.GET("/user/count", Handler.GetCountStudent)

		r.GET("/courses",Handler.GetAllMateri)
		r.GET("/course/id/:id",Handler.GetMateriByID)
		r.GET("/course/:page/:limit",Handler.GetMateriByPage)
		r.GET("/course/filter/:page/:limit",Handler.GetMateriByFilter)
		r.GET("/course/count",Handler.GetCountMateri)

		r.GET("/test/course/:id",Handler.GetTestByMateri)
		r.POST("/test/submit",Handler.SubmitTest)
		r.GET("/test/count",Handler.GetCountTest)
	}	

	r.Use(Handler.CheckAdmin)
	{
		r.POST("/course",Handler.CreateMateri)
		r.POST("/course/update",Handler.UpdateMateri)
		r.DELETE("/course/delete/:id",Handler.DeleteMateri)

		r.POST("/test",Handler.CreateTest)
		r.DELETE("/test/delete/:id",Handler.DeleteTest)
	}

	return r
}