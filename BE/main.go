package main

import (
	"belajar-golang/handler"
	"belajar-golang/repository"
	"belajar-golang/router"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main(){

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


	repo := repository.NewDB(db)
	Handler := handler.NewRepo(repo)
	r := router.Setuprouter(Handler)
	
	r.Run()
} 





