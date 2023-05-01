package models

type RefreshTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
