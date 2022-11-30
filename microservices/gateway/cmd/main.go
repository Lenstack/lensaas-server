package main

import (
	"github.com/Lenstack/Lensaas/tree/master/microservices/gateway/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()

	var (
		Environment = viper.Get("ENVIRONMENT").(string)
		GatewayPort = viper.Get("GATEWAY_PORT").(string)
	)

	loggerManager := infrastructure.NewLoggerManager(Environment)
	infrastructure.NewGraphqlServer(GatewayPort, loggerManager.Logger)
}
