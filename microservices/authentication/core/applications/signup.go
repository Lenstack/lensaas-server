package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignUp(ctx context.Context, req *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	token, err := ms.AuthenticationService.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pkg.SignUpResponse{Token: token}, nil
}
