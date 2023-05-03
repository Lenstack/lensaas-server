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
	authentication.HandleFunc("/sign_in", microservice.SignIn).Methods(http.MethodPost)                                    // Implemented
	authentication.HandleFunc("/sign_in_callback", microservice.SignInCallback).Methods(http.MethodGet)                    // Implemented
	authentication.HandleFunc("/sign_up", microservice.SignUp).Methods(http.MethodPost)                                    // Implemented
	authentication.HandleFunc("/sign_out", microservice.SignOut).Methods(http.MethodPost)                                  // Implemented
	authentication.HandleFunc("/refresh_token", microservice.RefreshToken).Methods(http.MethodPost)                        // Implemented
	authentication.HandleFunc("/forgot_password", microservice.ForgotPassword).Methods(http.MethodPost)                    // Not implemented
	authentication.HandleFunc("/reset_password", microservice.ResetPassword).Methods(http.MethodPost)                      // Not implemented
	authentication.HandleFunc("/verify_email", microservice.VerifyEmail).Methods(http.MethodGet)                           // Not implemented
	authentication.HandleFunc("/resend_email_verification", microservice.ResendEmailVerification).Methods(http.MethodPost) // Not implemented
	authentication.HandleFunc("/mfa", microservice.Mfa).Methods(http.MethodPost)                                           // Not implemented
	authentication.HandleFunc("/me", microservice.Me).Methods(http.MethodGet)                                              // Implemented
	return &Router{Handlers: router}
}
