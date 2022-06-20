package repository

import (
	"belajar-golang/model/user"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (h *Repository) GetCountStudent()(int,error){
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM users WHERE role = 'student' ").Scan(&count)
	if err != nil {
		return 0,err
	}
	return count, nil
}

func (h *Repository) Login(login user.Login) (user.User, error){
	sqlStatement := `SELECT id,name,role,email,password 
	FROM users WHERE email = ?`
	row := h.db.QueryRow(sqlStatement, login.Email)
	var user user.User
	err := row.Scan(&user.ID,&user.Name,&user.Role,&user.Email,&user.Password)
	if err != nil {
		return user,err
	}
	return user,nil
}

func (h *Repository) CheckEmailIsExist(email string) bool{
	sqlStatement := `SELECT id,name,role,email 
	FROM users WHERE email = ?`
	row := h.db.QueryRow(sqlStatement, email)
	var user user.User
	err := row.Scan(&user.ID,&user.Name,&user.Role,&user.Email)
	if err!=nil{
		return false
	}
	return true
}

func (h *Repository) GetUserByToken(token []string) (user.User, error){
	claims := &Claims{}
	var jwtKey = []byte("key")
	jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	sqlStatement := `SELECT id,name,role,email 
	FROM users WHERE email = ?`
	row := h.db.QueryRow(sqlStatement, claims.Email)
	var user user.User
	err := row.Scan(&user.ID,&user.Name,&user.Role,&user.Email)
	if err != nil{
		return user,err
	}
	return user,nil
}

func (h *Repository) CreateUser(user user.User) (user.User,error){

	sqlStmt := `INSERT INTO users (name, role,email,password,created_at, updated_at) 
	VALUES (?,?,?,?,?,?);`
	datetime := time.Now()  
	_, err := h.db.Exec(sqlStmt, user.Name,user.Role, user.Email, user.Password,datetime,datetime);
	if err != nil {
		return user, err
	}
 
	return user,nil
}




