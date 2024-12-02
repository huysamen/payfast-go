package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) CardTransactionQuery(idOrToken string) (rsp *types.Response[types.CreditCardStatus], errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/process/query/"+idOrToken, nil, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.Response[types.CreditCardStatus], types.ErrorResponse[int]](data, 200)
}
