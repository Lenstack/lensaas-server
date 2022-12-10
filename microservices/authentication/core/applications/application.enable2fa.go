package applications

import (
	"fmt"
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) Enable2FA(ctx context.Context, req *pkg.Enable2FARequest) (*pkg.Enable2FAResponse, error) {
	fmt.Println(req.GetPassword())
	return &pkg.Enable2FAResponse{Secret: "secret", QRCode: "qrcode"}, nil
}
