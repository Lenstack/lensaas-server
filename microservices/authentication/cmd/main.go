package main

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/applications"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/services"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/infrastructure"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/util"
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
		RedisHost                = viper.Get("REDIS_HOST").(string)
		RedisPort                = viper.Get("REDIS_PORT").(string)
		RedisPassword            = viper.Get("REDIS_PASSWORD").(string)
		JwtSecret                = viper.Get("JWT_SECRET").(string)
		JwtExpiration            = viper.Get("JWT_EXPIRATION").(string)
	)

	loggerManager := infrastructure.NewLoggerManager(Environment)
	postgres := infrastructure.NewPostgres(
		PostgresHost, PostgresPort, PostgresDatabaseName,
		PostgresDatabaseUser, PostgresDatabasePassword,
		loggerManager.Logger,
	)

	redis := infrastructure.NewRedisManager(RedisHost, RedisPort, RedisPassword, loggerManager.Logger)

	jwtManager := util.NewJwtManager(JwtSecret, JwtExpiration)

	authenticationService := services.NewAuthenticationService(postgres.Database, redis.Client, *jwtManager)
	middlewareApplication := applications.NewMiddlewareApplication(loggerManager.Logger)

	microservices := applications.NewMicroserviceServer(*middlewareApplication, *authenticationService)

	infrastructure.NewGrpcServer(GrpcPort, *microservices, loggerManager.Logger)
}
