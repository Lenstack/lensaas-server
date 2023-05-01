package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) ResetPassword(wr http.ResponseWriter, req *http.Request) {
	body := &models.ResetPasswordRequest{}
	// Decode reset password request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding sign in request"})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Validate sign up request body
	validateErrors := utils.Validate(body)
	if len(validateErrors) > 0 {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(validateErrors)
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Validate token and get user id
	userId, err := m.Jwt.ValidateToken(body.Token)
	if err != nil {
		m.Log.Error("Error validating token", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error validating token"})
		if err != nil {
			m.Log.Error("Error encoding reset password response", zap.Error(err))
		}
		return
	}

	// Reset password
	err = m.UserService.ResetPassword(userId, body.Password)
	if err != nil {
		m.Log.Error("Error resetting password", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error resetting password"})
		if err != nil {
			m.Log.Error("Error encoding reset password response", zap.Error(err))
		}
		return
	}

	// Encode reset password response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.ResetPasswordResponse{Code: http.StatusOK, Message: "Successfully reset password"})
	if err != nil {
		m.Log.Error("Error encoding reset password response", zap.Error(err))
	}
}
