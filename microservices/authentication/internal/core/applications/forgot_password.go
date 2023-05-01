package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) ForgotPassword(wr http.ResponseWriter, req *http.Request) {
	body := &models.ForgotPasswordRequest{}
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

	// Forgot password
	err = m.UserService.ForgotPassword(body.Email)
	if err != nil {
		m.Log.Error("Error forgotting password", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error forgotting password"})
		if err != nil {
			m.Log.Error("Error encoding forgot password response", zap.Error(err))
		}
		return
	}

	// Encode forgot password response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.ForgotPasswordResponse{Code: http.StatusOK, Message: "Successfully forgot password"})
	if err != nil {
		m.Log.Error("Error encoding forgot password response", zap.Error(err))
	}
}
