package main

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/setting/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()

	var (
		Environment              = viper.Get("ENVIRONMENT").(string)
		GrpcPort                 = viper.Get("GRPC_PORT").(string)
		PostgresHost             = viper.Get("POSTGRES_HOST").(string)
		PostgresPort             = viper.Get("POSTGRES_PORT").(string)
		PostgresDatabaseName     = viper.Get("POSTGRES_DATABASE_NAME").(string)
		PostgresDatabaseUser     = viper.Get("POSTGRES_DATABASE_USER").(string)
		PostgresDatabasePassword = viper.Get("POSTGRES_DATABASE_PASSWORD").(string)
	)

	loggerManager := infrastructure.NewLoggerManager(Environment)
	infrastructure.NewPostgres(
		PostgresHost, PostgresPort, PostgresDatabaseName,
		PostgresDatabaseUser, PostgresDatabasePassword,
		loggerManager.Logger,
	)

	infrastructure.NewGrpcServer(GrpcPort, loggerManager.Logger)
}
