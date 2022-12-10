package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type GraphqlServer struct {
}

func NewGraphqlServer(port string, router *chi.Mux, logger *zap.Logger) *GraphqlServer {
	logger.Sugar().Info("Graphql Server is running on port: ", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Sugar().Fatal("Failed to start Graphql Server: ", err)
		return nil
	}
	return &GraphqlServer{}
}
