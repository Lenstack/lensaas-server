package models

type CheckEmailAvailabilityRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type CheckEmailAvailabilityResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}
