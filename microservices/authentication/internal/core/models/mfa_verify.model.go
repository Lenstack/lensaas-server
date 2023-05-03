package models

type MFAVerifyRequest struct {
	UserId string `json:"user_id" validate:"required"`
	Code   string `json:"code" validate:"required"`
}

type MFAVerifyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
