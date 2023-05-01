package models

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type SignInResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SignInWithOAuthResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	RedirectUrl string `json:"redirect_url"`
}
