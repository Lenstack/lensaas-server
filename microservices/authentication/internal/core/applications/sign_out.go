package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) SignOut(wr http.ResponseWriter, req *http.Request) {
	// Create cookie store
	cookieStore := utils.NewCookieStore()

	// Get refresh token from cookie
	refreshToken, err := cookieStore.GetCookieDecrypted(req, cookieStore.RefreshTokenCookieName)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error getting refresh token from cookie"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Validate refresh token
	userId, err := m.Jwt.ValidateToken(refreshToken)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error validating refresh token"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Sign out user and revoke refresh token from database
	err = m.UserService.SignOut(userId, refreshToken)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Delete access token and refresh token from cookies
	cookieStore.DestroyCookie(wr, cookieStore.AccessTokenCookieName)
	cookieStore.DestroyCookie(wr, cookieStore.RefreshTokenCookieName)

	// Return response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.SignOutResponse{Code: http.StatusOK, Message: "Successfully signed out"})
	if err != nil {
		m.Log.Error("Error encoding sign in response", zap.Error(err))
	}
}
