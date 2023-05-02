package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) SignIn(wr http.ResponseWriter, req *http.Request) {
	// Start cookie store
	cookieStore := utils.NewCookieStore()
	// Get provider from query params
	provider := req.URL.Query().Get("provider")
	// If no provider != "" then credentials otherwise provider(social)
	if provider != "" {
		// Generate state
		codeGenerator := utils.NewCode()
		state, _ := codeGenerator.GenerateStateOauth()
		// Set state cookie
		cookieStore.SetCookieOAuth2State(wr, state)
		// Sign in with OAuth2 provider (social)
		url, err := m.UserService.SignInWithOAuth(provider, state)
		if err != nil {
			m.Log.Error("Error signing in with OAuth2", zap.Error(err))
			wr.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusInternalServerError, Message: err.Error()})
			if err != nil {
				m.Log.Error("Error encoding sign in response", zap.Error(err))
			}
			return
		}

		// Return redirect url
		wr.WriteHeader(http.StatusTemporaryRedirect)
		err = json.NewEncoder(wr).Encode(&models.SignInWithOAuthResponse{Code: http.StatusTemporaryRedirect, Message: "Signed in successfully with OAuth2", RedirectUrl: url})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Sign in with credentials (email and password)
	body := &models.SignInRequest{}
	// Decode sign in request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding sign in request"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Validate sign in request body
	validateErrors := utils.Validate(body)
	if len(validateErrors) > 0 {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(validateErrors)
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Sign in with credentials and get access and refresh tokens
	accessToken, refreshToken, err := m.UserService.SignInWithCredentials(body.Email, body.Password)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Session stuff here
	cookieStore.SetCookieAccessToken(wr, accessToken)
	cookieStore.SetCookieRefreshToken(wr, refreshToken)

	// Sign in response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.SignInResponse{Code: http.StatusOK, Message: "Signed in successfully with credentials"})
	if err != nil {
		m.Log.Error("Error encoding sign in response", zap.Error(err))
	}
}
