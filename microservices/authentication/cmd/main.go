package main

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/applications"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/services"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/infrastructure"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"github.com/spf13/viper"
)

func main() {
	// Load environment variables
	infrastructure.NewLoadEnv()
	var (
		AppEnvironment               = viper.GetString("APP_ENVIRONMENT")
		AppPort                      = viper.GetString("APP_PORT")
		JwtSecret                    = viper.GetString("JWT_SECRET")
		JwtAccessExpirationIn        = viper.GetString("JWT_ACCESS_EXPIRATION_IN")
		JwtRefreshExpirationIn       = viper.GetString("JWT_REFRESH_EXPIRATION_IN")
		JwtResetPasswordExpirationIn = viper.GetString("JWT_RESET_PASSWORD_EXPIRATION_IN")
	)

	// Initialize logger, database, jwt and other infrastructure
	logger := infrastructure.NewLogger(AppEnvironment)
	surrealDB := infrastructure.NewSurrealDB(logger)
	jwt := utils.NewJwt(JwtSecret, JwtAccessExpirationIn, JwtRefreshExpirationIn, JwtResetPasswordExpirationIn)
	bcrypt := utils.NewBcrypt()
	email := utils.NewEmail()

	// Initialize services (user, oauth provider, etc.)
	userService := services.NewUserService(surrealDB.Database, *jwt, *bcrypt, *email)

	// Initialize microservices
	microservices := applications.NewMicroservice(logger, *jwt, *userService)

	// Initialize router and start HTTP server
	router := infrastructure.NewRouter(*microservices)
	infrastructure.NewHttpServer(AppPort, router.Handlers, logger)
}
