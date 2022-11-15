package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignIn(ctx context.Context, req *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	token, err := ms.AuthenticationService.SignIn(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{Token: token}, nil
}
