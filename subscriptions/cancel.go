package subscriptions

import (
	"encoding/json"
	"strings"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Cancel(subscriptionID string) (bool, error) {
	body, err := c.put(strings.ReplaceAll(cancelPath, "__sid__", subscriptionID), nil)
	if err != nil {
		return false, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return false, err
	}

	if rsp.Code == 200 {
		return true, nil
	}

	return false, nil
}
