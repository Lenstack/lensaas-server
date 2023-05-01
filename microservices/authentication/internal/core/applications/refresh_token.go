package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) RefreshToken(wr http.ResponseWriter, req *http.Request) {
	cookieStore := utils.NewCookieStore()
	refreshToken, err := cookieStore.GetCookieDecrypted(req, cookieStore.RefreshTokenCookieName)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error getting refresh token from cookie"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	userId, err := m.Jwt.ValidateToken(refreshToken)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error validating refresh token"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	accessToken, err := m.UserService.RefreshToken(userId)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error refreshing token"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	cookieStore.SetCookieAccessToken(wr, accessToken)

	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.RefreshTokenResponse{Code: http.StatusOK, Message: "Successfully refreshed token"})
	if err != nil {
		m.Log.Error("Error encoding sign in response", zap.Error(err))
	}
}
