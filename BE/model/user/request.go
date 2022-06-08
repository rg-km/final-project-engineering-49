package user

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `binding:"required"`
}

type EditUserRequest struct {
	ID              int    `json:"id" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `binding:"required"`
}