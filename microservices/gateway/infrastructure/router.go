package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Router struct {
	App *chi.Mux
}

func NewRouter(routes []Route, logger *zap.Logger) *Router {
	app := chi.NewRouter()
	app.Use()

	for _, route := range routes {
		proxy, err := NewProxy(route.Target, route.Protocol, logger)
		if err != nil {
			logger.Sugar().Errorf("Error creating proxy: %v", err)
			return nil
		}
		app.Handle(route.Context, proxy)
	}
	return &Router{App: app}
}
