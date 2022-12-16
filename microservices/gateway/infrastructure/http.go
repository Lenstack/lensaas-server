package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type HttpTLS struct {
}

func NewHttpTLS(gatewayConfig Gateway, app *chi.Mux, logger *zap.Logger) *HttpTLS {
	listenServer, err := net.Listen("tcp", ":"+gatewayConfig.Listen.Port)
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
		return nil
	}

	logger.Sugar().Info("Graphql Server is running on port: ", gatewayConfig.Listen.Port)
	if err := http.ServeTLS(listenServer, app, "./cert/server-cert.pem", "./cert/server-key.pem"); err != nil {
		logger.Sugar().Fatal("Failed to start Graphql Server: ", err)
		return nil
	}
	return &HttpTLS{}
}
