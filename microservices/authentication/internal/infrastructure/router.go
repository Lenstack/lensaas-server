package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/applications"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Handlers http.Handler
}

func NewRouter(handlers *handler.Server, microservice applications.Microservice) *Router {
	router := mux.NewRouter()
	router.Use(microservice.MiddlewareCORS)
	router.Use(microservice.MiddlewareLogger)
	//router.Use(microservice.MiddlewareAuth)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handlers)

	return &Router{Handlers: router}
}
