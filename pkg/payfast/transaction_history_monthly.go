package payfast

import (
	"time"

	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

// TODO: determine unsuccessful response type
func (c *ClientImpl) TransactionHistoryMonthly(date *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error) {
	payload := new(specificDate)

	if date != nil {
		payload.Date = types.NewTime(*date)
	}

	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/transactions/history/monthly", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	txs, err := parseCsv(data)
	if err != nil {
		return nil, nil, err
	}

	return txs, nil, nil
}
