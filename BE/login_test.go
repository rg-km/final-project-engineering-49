package main

import (
	"net/http/httptest"
	"strings"

	"belajar-golang/handler"
	"belajar-golang/model/enrollment"
	"belajar-golang/model/lesson"
	"belajar-golang/model/log"
	"belajar-golang/model/materi"
	"belajar-golang/model/test"
	"belajar-golang/model/user"
	"belajar-golang/repository"
	"belajar-golang/router"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("JWT", func() {

	BeforeEach(func() {
		
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil{
			
		} 

		db.AutoMigrate(&user.User{}) 
		db.AutoMigrate(&enrollment.Enrollment{})
		db.AutoMigrate(&lesson.Lesson{})
		db.AutoMigrate(&log.Log{})
		db.AutoMigrate(&materi.Materi{})
		db.AutoMigrate(&test.Test{})  
	})

	Describe("/register", func() {
		It("should return successfully", func() {
			bodyReader := strings.NewReader(`
			{
				"Name" : "admin",
				"Email": "admin@gmail.com", 
				"Password": "admin123",
				"ConfirmPassword" : "admin123"
			}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
		
			req := httptest.NewRequest("POST" , "/register",bodyReader)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
		})
	})

	 
	Describe("/login", func() {
		It("should return non authorized when given wrong credential", func() {
			bodyReader := strings.NewReader(`{"email": "alamin@gmail.com", "password": "alamin"}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST" , "/login",bodyReader)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(400))
		})

		It("should return authorized when given correct credential", func() {
		
			bodyReader := strings.NewReader(`{"email": "admin@gmail.com", "password": "admin123"}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST" , "/login", bodyReader)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
		})
	})

	Describe("/materi", func() {

		It("create materi should return forbidden when request is not send token with admin auth", func() {	
			bodyReader := strings.NewReader(`
			{
				"Title" : "course 2",
				"Contain" : "desc course 2",
				"FileName": "course2.mp4",
				"Status": 1,
				"Creator": "Mr course2"
			}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST" , "/materi", bodyReader)
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(403))
		})
	})
})
