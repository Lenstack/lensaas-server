package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Lenstack/lensaas-server/microservices/authentication/graphql/directives"
	"github.com/Lenstack/lensaas-server/microservices/authentication/graphql/generated"
	"github.com/Lenstack/lensaas-server/microservices/authentication/graphql/resolvers"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/applications"
)

type GraphqlServer struct {
	Handlers *handler.Server
}

func NewGraphqlServer(microservice applications.Microservice) *GraphqlServer {
	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			Microservice: microservice,
		},
		Directives: generated.DirectiveRoot{
			HasAuth:       directives.HasAuth,
			HasPermission: directives.HasPermission,
		},
	}))
	return &GraphqlServer{Handlers: graphqlServer}
}
