package auth

type UserLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserAuthorization struct {
	UserLogin
	Email string `json:"email" binding:"required,email"`
}
