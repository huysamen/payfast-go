package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryMonthlyReq struct {
	Date types.Time `payfast:"date,query,yyyy-mm-dd,optional"`
}

func (c *Client) Monthly(payload TransactionHistoryMonthlyReq) (txs []*types.Transaction, status int, err error) {
	body, status, err := c.get(monthlyPath, payload)
	if err != nil {
		return nil, status, err
	}

	txs, err = parseCsv(body)
	if err != nil {
		return nil, status, err
	}

	return txs, status, nil
}
