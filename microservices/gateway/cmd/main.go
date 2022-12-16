package main

import (
	"github.com/Lenstack/Lensaas/microservices/gateway/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()

	gatewayConfig := &infrastructure.Gateway{}
	err := viper.UnmarshalKey("gateway", gatewayConfig)
	if err != nil {
		panic(err)
	}

	loggerManager := infrastructure.NewLoggerManager(gatewayConfig.Environment)
	graphqlServer := infrastructure.NewGraphqlServer(gatewayConfig.Services)
	graphqlServer.CallGrpcMicroservices()

	app := infrastructure.NewRouter(graphqlServer.GraphqlServer)
	infrastructure.NewHttpTLS(*gatewayConfig, app.Router, loggerManager.Logger)
}
