package main

import (
	"github.com/Lenstack/Lensaas/microservices/gateway/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()

	var (
		environment = viper.GetString("gateway.environment")
	)

	loggerManager := infrastructure.NewLoggerManager(environment)

	gatewayConfig := &infrastructure.GatewayConfig{}
	err := viper.UnmarshalKey("gateway", gatewayConfig)
	if err != nil {
		panic(err)
	}

	router := infrastructure.NewRouter(gatewayConfig.Routes, loggerManager.Logger)
	infrastructure.NewGraphqlServer(gatewayConfig.Listen.Port, router.App, loggerManager.Logger)
}
