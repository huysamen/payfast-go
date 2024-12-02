package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) SubscriptionFetch(token string) (rsp *types.Response[types.Subscription], errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/subscriptions/"+token+"/fetch", nil, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.Response[types.Subscription], types.ErrorResponse[int]](data, 200)
}
