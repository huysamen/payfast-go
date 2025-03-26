package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryWeeklyReq struct {
	Date types.Time `payfast:"date,query,yyyy-mm-dd,optional"`
}

func (c *Client) Weekly(payload TransactionHistoryWeeklyReq) (txs []*types.Transaction, status int, err error) {
	body, status, err := c.get(weeklyPath, payload)
	if err != nil {
		return nil, status, err
	}

	txs, err = parseCsv(body)
	if err != nil {
		return nil, status, err
	}

	return txs, status, nil
}
