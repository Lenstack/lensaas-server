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
	router.Use(microservice.MiddlewareCORS)
	router.Use(microservice.MiddlewareCleanPath)
	router.Use(microservice.MiddlewareAllowContentType)

	public := router.PathPrefix("/v1/").Subrouter()
	public.HandleFunc("/sign_in", microservice.SignIn).Methods("POST")
	public.HandleFunc("/sign_in_callback", microservice.SignInCallback).Methods("POST")
	public.HandleFunc("/sign_up", microservice.SignUp).Methods("POST")
	public.HandleFunc("/sign_out", microservice.SignOut).Methods("POST")
	public.HandleFunc("/forgot_password", microservice.ForgotPassword).Methods("POST")
	public.HandleFunc("/reset_password", microservice.ResetPassword).Methods("POST")
	public.HandleFunc("/send_verification_email", microservice.SendVerificationEmail).Methods("POST")
	public.HandleFunc("/verify_email", microservice.VerifyEmail).Methods("POST")
	public.HandleFunc("/check_email_availability", microservice.CheckEmailAvailability).Methods("POST")
	public.HandleFunc("/refresh_token", microservice.RefreshToken).Methods("POST")
	public.HandleFunc("/mfa_enable", microservice.MFAEnable).Methods("POST")
	public.HandleFunc("/mfa_disable", microservice.MFADisable).Methods("POST")

	private := router.PathPrefix("/v1/").Subrouter()
	private.Use(microservice.MiddlewareAuthentication)
	private.HandleFunc("/me", microservice.Me).Methods("GET")
	private.HandleFunc("/ai_assistant", microservice.AIAssistant).Methods("POST")
	private.HandleFunc("/users", microservice.AIAssistant).Methods("GET")
	private.HandleFunc("/users/{id}", microservice.AIAssistant).Methods("GET")

	return &Router{Handlers: router}
}
