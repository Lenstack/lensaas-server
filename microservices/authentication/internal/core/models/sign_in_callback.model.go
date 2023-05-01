package models

type SignInCallbackResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
