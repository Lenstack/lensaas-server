package models

type MFAGenerateRequest struct {
	UserId string `json:"user_id"`
}

type MFAGenerateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
	QRCode  string `json:"qr_code"`
}
