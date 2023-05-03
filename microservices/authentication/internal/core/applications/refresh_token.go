package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) RefreshToken(wr http.ResponseWriter, req *http.Request) {
	// Initialize cookie store
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

	// Refresh token and get new access token
	accessToken, err := m.UserService.RefreshToken(userId, refreshToken)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Set access token in cookie
	cookieStore.SetCookieAccessToken(wr, accessToken)

	// Return response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.RefreshTokenResponse{Code: http.StatusOK, Message: "Successfully refreshed token"})
	if err != nil {
		m.Log.Error("Error encoding sign in response", zap.Error(err))
	}
}
