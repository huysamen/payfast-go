package subscriptions

import (
	"encoding/json"
	"strings"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Fetch(token string) (*types.Subscription, error) {
	body, err := c.get(strings.ReplaceAll(fetchPath, "__token__", token))
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
		sub.Copy(rsp.Data.Response.(map[string]interface{}))

		return sub, nil
	}

	return nil, nil
}
