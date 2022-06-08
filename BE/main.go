package main

import (
	"belajar-golang/handler"
	"belajar-golang/model/enrollment"
	"belajar-golang/model/lesson"
	"belajar-golang/model/log"
	"belajar-golang/model/materi"
	"belajar-golang/model/test"
	"belajar-golang/model/user"
	"belajar-golang/repository"
	"belajar-golang/router"
	"time"

	"github.com/gin-contrib/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 

var db *gorm.DB

func main(){
	
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{} 

	db.AutoMigrate(&user.User{}) 
	db.AutoMigrate(&materi.Materi{})
	db.AutoMigrate(&test.Test{})  
	
	db.AutoMigrate(&enrollment.Enrollment{})
	db.AutoMigrate(&lesson.Lesson{})
	db.AutoMigrate(&log.Log{})
	   

	repo := repository.NewDB(db)
	Handler := handler.NewRepo(repo)
	r := router.Setuprouter(Handler)

 
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"PUT","PATCH","GET","POST"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))

	r.Run()
} 





