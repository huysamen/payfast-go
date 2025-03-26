package subscriptions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Cancel(token string) (ok bool, status int, err error) {
	body, status, err := c.put(PathCat(basePath, token, cancelPath), nil)
	if err != nil {
		return false, status, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return false, status, err
	}

	if rsp.Code == 200 {
		return true, status, nil
	}

	return false, status, nil
}
