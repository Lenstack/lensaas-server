package models

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type SignUpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
