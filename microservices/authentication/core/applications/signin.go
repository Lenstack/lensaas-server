package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignIn(ctx context.Context, req *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	accessToken, expiration, err := ms.AuthenticationService.SignIn(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{AccessToken: accessToken, Expiration: expiration}, nil
}
