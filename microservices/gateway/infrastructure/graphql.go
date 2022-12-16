package infrastructure

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Lenstack/Lensaas/microservices/gateway/graphql/generated"
	"github.com/Lenstack/Lensaas/microservices/gateway/graphql/resolvers"
)

type GraphqlServer struct {
	GraphqlServer *handler.Server
	services      []Service
}

func NewGraphqlServer(services []Service) *GraphqlServer {
	graphqlServer := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolvers.Resolver{}},
		),
	)
	return &GraphqlServer{GraphqlServer: graphqlServer, services: services}
}

func (gs *GraphqlServer) CallGrpcMicroservices() {
	for _, service := range gs.services {
		fmt.Println(service)
	}
}
