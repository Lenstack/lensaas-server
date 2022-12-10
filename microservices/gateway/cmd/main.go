package main

import (
	"github.com/Lenstack/Lensaas/microservices/gateway/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()
	loggerManager := infrastructure.NewLoggerManager("development")

	gatewayConfig := &infrastructure.GatewayConfig{}
	err := viper.UnmarshalKey("gateway", gatewayConfig)
	if err != nil {
		panic(err)
	}

	router := infrastructure.NewRouter(gatewayConfig.Routes, loggerManager.Logger)
	infrastructure.NewGraphqlServer(gatewayConfig.Listen.Port, router.App, loggerManager.Logger)
}
