package subscriptions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Fetch(token string) (subscription *types.Subscription, status int, err error) {
	body, status, err := c.get(PathCat(basePath, token, fetchPath), nil)
	if err != nil {
		return nil, status, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return nil, status, err
	}

	if rsp.Code == 200 {
		subscription = &types.Subscription{}
		subscription.Copy(rsp.Data.Response.(map[string]any))

		return subscription, status, nil
	}

	return nil, status, nil
}
