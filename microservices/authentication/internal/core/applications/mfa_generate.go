package applications

import (
	"encoding/json"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/models"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (m *Microservice) MFAGenerate(wr http.ResponseWriter, req *http.Request) {
	body := &models.MFAGenerateRequest{}
	// Decode mfa generate request body
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: "Error decoding mfa generate request"})
		if err != nil {
			m.Log.Error("Error encoding mfa generate response", zap.Error(err))
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

	// MFA Generate
	qrCode, token, err := m.UserService.MFAGenerate(body.UserId)
	if err != nil {
		m.Log.Error("Error mfa generate", zap.Error(err))
		wr.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(wr).Encode(&models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			m.Log.Error("Error encoding mfa generate response", zap.Error(err))
		}
		return
	}

	// Encode mfa generate response
	wr.WriteHeader(http.StatusOK)
	err = json.NewEncoder(wr).Encode(&models.MFAGenerateResponse{Code: http.StatusOK, Message: "Successfully mfa generate",
		QRCode: qrCode, Token: token})
	if err != nil {
		m.Log.Error("Error encoding mfa generate response", zap.Error(err))
	}

}
