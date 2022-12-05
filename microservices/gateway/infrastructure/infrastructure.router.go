package infrastructure

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"moul.io/chizap"
)

type Router struct {
	app *chi.Mux
}

func NewRouter(srv *handler.Server, logger *zap.Logger) *Router {
	app := chi.NewRouter()
	app.Use(cors.Handler(
		cors.Options{
			AllowCredentials: true,
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		},
	))
	app.Use(chizap.New(logger, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))

	app.Handle("/", playground.Handler("GraphQL playground", "/query"))
	app.Handle("/query", srv)
	return &Router{app: app}
}
