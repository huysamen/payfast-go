package payfast

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) RefundQuery(pfPaymentID string) (rsp *types.RefundQuery, errRsp *types.ErrorResponse[int], err error) {
	data, status, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/refunds/query/"+pfPaymentID, nil, c.testing)
	if err != nil {
		return nil, nil, err
	}

	// todo: not sure how this api responds to non-ok scenarios
	if status != 200 {
		return nil, nil, nil
	}

	rsp = new(types.RefundQuery)

	err = json.Unmarshal(data, rsp)
	if err != nil {
		return nil, nil, err
	}

	return rsp, nil, nil
}
