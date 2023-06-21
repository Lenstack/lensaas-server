package models

type SendVerificationEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SendVerificationEmailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
