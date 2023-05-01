package models

type ResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type ResetPasswordResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
