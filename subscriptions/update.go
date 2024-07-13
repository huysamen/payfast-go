package subscriptions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

type UpdateSubscriptionReq struct {
	Cycles    types.Numeric `payfast:"cycles,body,numeric,optional"`      // The number of cycles for the subscription.
	Frequency types.Numeric `payfast:"frequency,body,numeric,optional"`   // The frequency for the subscription
	RunDate   types.Time    `payfast:"run_date,body,yyyy-mm-dd,optional"` // The next run date for the subscription. YYYY-MM-DD
	Amount    types.Numeric `payfast:"amount,body,numeric,optional"`      // The amount which the buyer must pay, in CENTS (ZAR).
}

func (c *Client) Update(token string, payload UpdateSubscriptionReq) (*types.Subscription, error) {
	body, err := c.patch(PathCat(basePath, token, updatePath), payload)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return nil, err
	}

	if rsp.Code == 200 {
		sub := &types.Subscription{}
		sub.Copy(rsp.Data.Response.(map[string]any))

		return sub, nil
	}

	return nil, nil
}
