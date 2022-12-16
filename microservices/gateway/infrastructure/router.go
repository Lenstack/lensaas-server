package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter(graphqlServer *handler.Server) *Router {
	app := chi.NewRouter()
	app.Use(cors.Handler(
		cors.Options{
			AllowCredentials: true,
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		},
	))
	app.Handle("/", playground.Handler("GraphQL playground", "/query"))
	app.Handle("/query", graphqlServer)
	return &Router{Router: app}
}
