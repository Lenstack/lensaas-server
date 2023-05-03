package models

type MFADisableRequest struct {
	UserId string `json:"user_id" validate:"required"`
}

type MFADisableResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
