package models

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email string `json:"email" biding:"required,email"`
	Password string `json:"password" biding:"required,min=8"`
}

type LoginRequest struct {
	Email string `json:"email" biding:"required,email"`
	Password string `json:"password" biding:"required"`
}