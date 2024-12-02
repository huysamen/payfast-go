package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) SubscriptionUnpause(token string) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Put(c.client, c.merchantID, c.merchantPassphrase, "/subscriptions/"+token+"/unpause", nil, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.ConfirmationResponse, types.ErrorResponse[int]](data, 200)
}
