package models

type MeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    UserMe `json:"user"`
}

type UserMe struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
