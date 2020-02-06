package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryMonthlyReq struct {
	Date types.Time `payfast:"date,query,yyyy-mm-dd,optional"`
}

func (c *Client) Monthly(payload TransactionHistoryMonthlyReq) ([]*types.Transaction, error) {
	body, err := c.get(monthlyPath, payload)
	if err != nil {
		return nil, err
	}

	txs, err := parseCsv(body)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
