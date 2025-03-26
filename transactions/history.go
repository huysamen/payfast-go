package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryReq struct {
	From types.Time `payfast:"from,query,yyyy-mm-dd,optional"`
	To   types.Time `payfast:"to,query,yyyy-mm-dd,optional"`
}

func (c *Client) History(payload TransactionHistoryReq) (txs []*types.Transaction, status int, err error) {
	body, status, err := c.get(historyPath, payload)
	if err != nil {
		return nil, status, err
	}

	txs, err = parseCsv(body)
	if err != nil {
		return nil, status, err
	}

	return txs, status, nil
}
