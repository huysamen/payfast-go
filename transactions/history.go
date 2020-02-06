package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

type TransactionHistoryReq struct {
	From types.Time `payfast:"from,query,yyyy-mm-dd,optional"`
	To   types.Time `payfast:"to,query,yyyy-mm-dd,optional"`
}

func (c *Client) History(payload TransactionHistoryReq) ([]*types.Transaction, error) {
	body, err := c.get(historyPath, payload)
	if err != nil {
		return nil, err
	}

	txs, err := parseCsv(body)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
