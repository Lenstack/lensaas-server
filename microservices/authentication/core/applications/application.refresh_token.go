package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
)

func (ms *MicroserviceServer) RefreshToken(req *pkg.RefreshTokenRequest, stream pkg.Authentication_RefreshTokenServer) error {
	err := stream.Send(&pkg.RefreshTokenResponse{AccessToken: "new access token", Expiration: "1234567890"})
	if err != nil {
		return err
	}
	return nil
}
