package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) MFAVerify(wr http.ResponseWriter, req *http.Request) {
	body := &models.MFAVerifyRequest{}
	// Decode mfa verify request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding mfa verify request"})
		if err != nil {
			m.Log.Error("Error encoding mfa verify response", zap.Error(err))
		}
		return
	}

	// Validate mfa verify request body
	validateErrors := utils.Validate(body)
	if len(validateErrors) > 0 {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(validateErrors)
		if err != nil {
			m.Log.Error("Error encoding mfa verify response", zap.Error(err))
		}
		return
	}

	// MFA Verify

	// Encode mfa verify response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.MFAVerifyResponse{Code: http.StatusOK, Message: "Successfully mfa verify"})
	if err != nil {
		m.Log.Error("Error encoding mfa verify response", zap.Error(err))
	}
}
