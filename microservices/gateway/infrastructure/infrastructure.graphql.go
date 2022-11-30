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
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)

	logger.Sugar().Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Sugar().Fatalf("failed to start server: %v", err)
		return nil
	}
	return &GraphqlServer{}
}
