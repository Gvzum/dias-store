package auth

type UserSignInSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserSignUpSchema struct {
	UserSignInSchema
	Name string `json:"name" binding:"required"`
}
