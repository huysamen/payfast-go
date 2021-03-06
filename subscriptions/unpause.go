package subscriptions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Unpause(token string) (bool, error) {
	body, err := c.put(PathCat(basePath, token, unpausePath), nil)
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
