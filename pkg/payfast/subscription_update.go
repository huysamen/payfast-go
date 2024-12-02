package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type SubscriptionUpdateRequest struct {
	Cycles    types.Numeric
	Frequency types.Numeric
	RunDate   types.Time
	Amount    types.Numeric
}

func (r *SubscriptionUpdateRequest) Headers() map[string]string {
	return map[string]string{}
}

func (r *SubscriptionUpdateRequest) Query() map[string]string {
	return map[string]string{}
}

func (r *SubscriptionUpdateRequest) Body() map[string]net.BodyData {
	return map[string]net.BodyData{
		"cycles":    r.Cycles,
		"frequency": r.Frequency,
		"run_date":  r.RunDate,
		"amount":    r.Amount,
	}
}

// TODO: determine unsuccessful response type
func (c *ClientImpl) SubscriptionUpdate(
	token string,
	payload *SubscriptionUpdateRequest,
) (
	rsp *types.Response[types.Subscription],
	errRsp *types.ErrorResponse[int],
	err error,
) {
	data, _, err := net.Post(c.client, c.merchantID, c.merchantPassphrase, "/subscriptions/"+token+"/update", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.Response[types.Subscription], types.ErrorResponse[int]](data, 200)
}
