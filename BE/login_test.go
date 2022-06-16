package main

import (
	"belajar-golang/handler"
	"belajar-golang/repository"
	"belajar-golang/router"
	"database/sql"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWT", func() {

	BeforeEach(func() {
		
	var err error	
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil{
		panic(err)
	}
	
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS materis (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			title VARCHAR(255) NOT NULL,
			contain TEXT NOT NULL,
			creator VARCHAR(255) NOT NULL,
			file_name VARCHAR(255),
    		created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
	);`)
	if err != nil {
		panic(err)
	}


	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			name VARCHAR(255) NOT NULL,
			role VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
    		created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
	);`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			user_id INTEGER NOT NULL,
			materi_id INTEGER NOT NULL,
			score INTEGER NOT NULL,
    		created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
	);`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tests (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			materi_id INTEGER NOT NULL,
			question VARCHAR(255) NOT NULL,
			answer1 VARCHAR(255) NOT NULL,
			answer2 VARCHAR(255) NOT NULL,
			answer3 VARCHAR(255) NOT NULL,
			answer4 VARCHAR(255) NOT NULL,
			answer_true VARCHAR(255) NOT NULL,
    		created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
	);`)
	if err != nil {
		panic(err)
	}
		
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

})
