package models

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
