package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) DeleteAccount(ctx context.Context, req *pkg.DeleteAccountRequest) (*pkg.DeleteAccountResponse, error) {
	return &pkg.DeleteAccountResponse{Message: "the account has been deleted"}, nil
}
