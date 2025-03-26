package transactions

import (
	"encoding/json"

	"github.com/huysamen/payfast-go/types"
)

func (c *Client) Query(token string) (cs *types.CreditCardStatus, status int, err error) {
	body, status, err := c.get(PathCat(queryPath, token), nil)
	if err != nil {
		return nil, status, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return nil, status, err
	}

	if rsp.Code == 200 {
		cs := &types.CreditCardStatus{}
		cs.Copy(rsp.Data.Response.(map[string]any))

		return cs, status, nil
	}

	return nil, status, nil
}
