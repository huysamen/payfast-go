package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) RefundRetrieve(pfPaymentID string) (rsp *types.Response[types.Refund], errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/refunds/"+pfPaymentID, nil, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.Response[types.Refund], types.ErrorResponse[int]](data, 200)
}
