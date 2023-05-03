package models

type ResendEmailVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResendEmailVerificationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
