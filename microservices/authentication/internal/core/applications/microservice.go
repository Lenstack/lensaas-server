package applications

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/services"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"go.uber.org/zap"
)

type Microservice struct {
	Log         *zap.Logger
	Jwt         utils.Jwt
	UserService services.UserService
}

func NewMicroservice(log *zap.Logger, jwt utils.Jwt, userService services.UserService) *Microservice {
	return &Microservice{
		Log:         log,
		Jwt:         jwt,
		UserService: userService,
	}
}
