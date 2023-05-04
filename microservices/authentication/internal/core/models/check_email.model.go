package models

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type CheckEmailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}
