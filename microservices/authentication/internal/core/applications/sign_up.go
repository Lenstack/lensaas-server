package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) SignUp(wr http.ResponseWriter, req *http.Request) {
	// Create sign up request body
	body := &models.SignUpRequest{}
	// Decode sign up request body
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

	// Sign up with credentials (email and password)
	err = m.UserService.SignUp(body.Name, body.Email, body.Password)
	if err != nil {
		m.Log.Error("Error signing up with credentials", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding sign in response", zap.Error(err))
		}
		return
	}

	// Encode sign up response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.SignUpResponse{Code: http.StatusOK, Message: "Successfully signed up"})
	if err != nil {
		m.Log.Error("Error encoding sign up response", zap.Error(err))
	}
}
