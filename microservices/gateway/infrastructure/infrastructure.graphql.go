package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Lenstack/Lensaas/tree/master/microservices/gateway/graph"
	"github.com/Lenstack/Lensaas/tree/master/microservices/gateway/graph/generated"
	"go.uber.org/zap"
	"net/http"
)

type GraphqlServer struct {
	port string
}

func NewGraphqlServer(port string, logger *zap.Logger) *GraphqlServer {
	graphqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router := NewRouter(graphqlSrv, logger)

	logger.Sugar().Info("Graphql Server is running on port: ", port)
	if err := http.ListenAndServe(":"+port, router.app); err != nil {
		logger.Sugar().Fatal("Failed to start Graphql Server: ", err)
		return nil
	}
	return &GraphqlServer{}
}
