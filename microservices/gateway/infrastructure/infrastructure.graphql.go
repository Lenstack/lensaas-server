package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Lenstack/Lensaas/tree/master/microservices/gateway/graph"
	"github.com/Lenstack/Lensaas/tree/master/microservices/gateway/graph/generated"
	"go.uber.org/zap"
	"net/http"
)

type GraphqlServer struct {
	port string
}

func NewGraphqlServer(port string, logger *zap.Logger) *GraphqlServer {

	// Create a new GraphqlServer instance
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Sugar().Info("GraphqlServer is running on port: ", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Sugar().Fatal("Failed to start GraphqlServer: ", err)
		return nil
	}
	return &GraphqlServer{}
}
