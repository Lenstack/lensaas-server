package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) MFAEnable(wr http.ResponseWriter, req *http.Request) {
	body := &models.MFAEnableRequest{}
	// Decode mfa enable request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding mfa enable request"})
		if err != nil {
			m.Log.Error("Error encoding mfa enable response", zap.Error(err))
		}
		return
	}

	// Validate mfa enable request body
	validateErrors := utils.Validate(body)
	if len(validateErrors) > 0 {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(validateErrors)
		if err != nil {
			m.Log.Error("Error encoding mfa enable response", zap.Error(err))
		}
		return
	}

	// MFA Enable

	// Encode mfa enable response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.MFAEnableResponse{Code: http.StatusOK, Message: "Successfully mfa enable"})
	if err != nil {
		m.Log.Error("Error encoding mfa enable response", zap.Error(err))
	}
}
