package main

import "github.com/Lenstack/Lensaas/tree/master/microservices/gateway/infrastructure"

func main() {
	loggerManager := infrastructure.NewLoggerManager("development")
	infrastructure.NewGraphqlServer("8080", loggerManager.Logger)
}
