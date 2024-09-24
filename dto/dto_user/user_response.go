package dto

type UserRegisterResponse struct {
	ID       uint   `json:"id"  example:"1"`
	Username string `json:"username"  example:"username"`
	Email    string `json:"email" example:"username@email.com"`
	Age      uint   `json:"age" example:"11"`
}

type LoginResponse struct {
	User  UserRegisterResponse
	Token string `json:"token"`
}
