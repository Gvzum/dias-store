package auth

type UserSignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserSignUp struct {
	UserSignIn
	Name string `json:"name" binding:"required"`
}
