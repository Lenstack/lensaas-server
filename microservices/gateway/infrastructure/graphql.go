package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type GraphqlServer struct {
}

func NewGraphqlServer(port string, router *chi.Mux, logger *zap.Logger) *GraphqlServer {
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
		return nil
	}

	logger.Sugar().Info("Graphql Server is running on port: ", port)
	if err := http.ServeTLS(listenServer, router, "./cert/server-cert.pem", "./cert/server-key.pem"); err != nil {
		logger.Sugar().Fatal("Failed to start Graphql Server: ", err)
		return nil
	}
	return &GraphqlServer{}
}
