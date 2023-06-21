package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) CheckEmailAvailability(wr http.ResponseWriter, req *http.Request) {
	body := &models.CheckEmailAvailabilityRequest{}
	// Decode check email request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding check email request"})
		if err != nil {
			m.Log.Error("Error encoding check email response", zap.Error(err))
		}
		return
	}

	// Validate check email request body
	validateErrors := utils.Validate(body)
	if len(validateErrors) > 0 {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(validateErrors)
		if err != nil {
			m.Log.Error("Error encoding check email response", zap.Error(err))
		}
		return
	}

	// Check email
	err = m.UserService.CheckEmail(body.Email)
	if err != nil {
		m.Log.Error("Error checking email", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding check email response", zap.Error(err))
		}
		return
	}

	// Encode check email response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.CheckEmailAvailabilityResponse{Code: http.StatusOK, Message: "Successfully checked email"})
	if err != nil {
		m.Log.Error("Error encoding check email response", zap.Error(err))
	}
}
