package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) SignInCallback(wr http.ResponseWriter, req *http.Request) {
	// Get provider from query params
	provider := req.URL.Query().Get("provider")

	// If no provider != "" then error
	if provider == "" {
		wr.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusBadRequest, Message: "Invalid provider from query params"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Start cookie store
	cookieStore := utils.NewCookieStore()

	// Get state from cookie
	oauthState, err := cookieStore.GetCookieDecrypted(req, cookieStore.OauthStateCookieName)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusInternalServerError, Message: "Error getting state from cookie"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Get state and code from query params
	state := req.FormValue("state")
	code := req.FormValue("code")

	// Validate state from cookie and state from query params
	if state != oauthState {
		wr.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusInternalServerError, Message: "Invalid state from cookie and state from query params"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Sign in with OAuth2 provider (social)
	accessToken, refreshToken, err := m.UserService.SignInWithOAuthCallback(provider, code)
	if err != nil {
		m.Log.Error("Error signing in with OAuth2", zap.Error(err))
		wr.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Set access token cookie
	cookieStore.SetCookieAccessToken(wr, accessToken)
	cookieStore.SetCookieRefreshToken(wr, refreshToken)

	// Redirect to home page
	wr.WriteHeader(http.StatusTemporaryRedirect)
	http.Redirect(wr, req, "/", http.StatusTemporaryRedirect)
}
