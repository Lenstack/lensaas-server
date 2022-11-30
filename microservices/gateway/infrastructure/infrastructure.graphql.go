package infrastructure

import (
	"go.uber.org/zap"
	"net/http"
)

type GraphqlServer struct {
	port string
}

func NewGraphqlServer(port string, logger *zap.Logger) *GraphqlServer {

	// Create a new GraphqlServer instance

	logger.Sugar().Info("GraphqlServer is running on port: ", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Sugar().Fatal("Failed to start GraphqlServer: ", err)
		return nil
	}
	return &GraphqlServer{}
}
