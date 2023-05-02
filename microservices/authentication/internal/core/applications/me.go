package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) Me(wr http.ResponseWriter, req *http.Request) {
	cookieStore := utils.NewCookieStore()
	accessToken, err := cookieStore.GetCookieDecrypted(req, cookieStore.AccessTokenCookieName)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error getting access token from cookie"})
		return
	}

	userId, err := m.Jwt.ValidateToken(accessToken)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error validating access token"})
		return
	}

	user, err := m.UserService.Me(userId)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.MeResponse{Code: http.StatusOK, Message: "Successfully retrieved user",
		User: models.UserMe{
			Id:     user.ID,
			Email:  user.Email,
			Name:   user.Profile.Name,
			Avatar: user.Profile.Avatar,
		}})
	if err != nil {
		m.Log.Error("Error encoding sign in response", zap.Error(err))
	}
}
