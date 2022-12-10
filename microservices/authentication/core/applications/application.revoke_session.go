package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) RevokeSession(ctx context.Context, req *pkg.RevokeSessionRequest) (*pkg.RevokeSessionResponse, error) {
	return &pkg.RevokeSessionResponse{Message: "the session has been revoked"}, nil
}
