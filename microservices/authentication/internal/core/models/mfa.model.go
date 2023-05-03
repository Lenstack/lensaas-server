package models

type MfaRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type MfaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
