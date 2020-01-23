package subscriptions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

type PauseSubscriptionReq struct {
	Cycles types.Numeric `payfast:"cycles,body,numeric,optional"`
}

func (c *Client) Pause(token string, payload PauseSubscriptionReq) (bool, error) {
	body, err := c.put(PathCat(basePath, token, pausePath), payload)
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
