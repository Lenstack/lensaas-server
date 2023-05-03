package models

type ResendEmailVerificationRequest struct {
	Email string `json:"email"`
}

type ResendEmailVerificationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
