package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryWeeklyReq struct {
	Date types.Time `payfast:"date,query,yyyy-mm-dd,optional"`
}

func (c *Client) Weekly(payload TransactionHistoryWeeklyReq) ([]*types.Transaction, error) {
	body, err := c.get(weeklyPath, payload)
	if err != nil {
		return nil, err
	}

	txs, err := parseCsv(body)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
