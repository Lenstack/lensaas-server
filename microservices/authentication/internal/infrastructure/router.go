package infrastructure

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/applications"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Handlers http.Handler
}

func NewRouter(microservice applications.Microservice) *Router {
	router := mux.NewRouter()
	router.Use(microservice.MiddlewareCleanPath)
	router.Use(microservice.MiddlewareCORS)
	//router.Use(microservice.MiddlewareCsrf)
	router.Use(microservice.MiddlewareLogger)
	router.Use(microservice.MiddlewareRecovery)
	router.Use(microservice.MiddlewareRateLimit)

	authentication := router.PathPrefix("/v1/authentication").Subrouter()
	authentication.HandleFunc("/sign_in", microservice.SignIn).Methods(http.MethodPost)
	authentication.HandleFunc("/sign_in_callback", microservice.SignInCallback).Methods(http.MethodGet)
	authentication.HandleFunc("/sign_up", microservice.SignUp).Methods(http.MethodPost)
	authentication.HandleFunc("/sign_out", microservice.SignOut).Methods(http.MethodPost)
	authentication.HandleFunc("/refresh_token", microservice.RefreshToken).Methods(http.MethodPost)

	authentication.HandleFunc("/forgot_password", microservice.ForgotPassword).Methods(http.MethodPost)
	authentication.HandleFunc("/reset_password", microservice.ResetPassword).Methods(http.MethodPost)
	authentication.HandleFunc("/verify_email", microservice.VerifyEmail).Methods(http.MethodGet)
	authentication.HandleFunc("/resend_email_verification", microservice.ResendEmailVerification).Methods(http.MethodPost)

	authentication.HandleFunc("/mfa_enable", microservice.MFAEnable).Methods(http.MethodPost)
	authentication.HandleFunc("/mfa_disable", microservice.MFADisable).Methods(http.MethodPost)
	authentication.HandleFunc("/mfa_verify", microservice.MFAVerify).Methods(http.MethodPost)

	authentication.HandleFunc("/me", microservice.Me).Methods(http.MethodGet)

	return &Router{Handlers: router}
}
