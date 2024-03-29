package models

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token"`
}

type VerifyEmailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
