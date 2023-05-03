package models

type MFAEnableRequest struct {
	UserId string `json:"user_id" validate:"required"`
}

type MFAEnableResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	QRCode  string `json:"qr_code"`
}
