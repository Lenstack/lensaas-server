package infrastructure

import (
	"go.uber.org/zap"
	"net"
	"net/http"
)

type HttpServer struct {
	Port string
}

func NewHttpServer(port string, handlers http.Handler, log *zap.Logger) *HttpServer {
	log.Info("Starting http server")
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to start http server", zap.Error(err))
	}

	log.Info("Started http server on port " + port)
	err = http.Serve(listen, handlers)
	if err != nil {
		log.Fatal("failed to start http server", zap.Error(err))
	}
	return &HttpServer{port}
}
