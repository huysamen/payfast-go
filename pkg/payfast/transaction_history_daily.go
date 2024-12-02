package payfast

import (
	"time"

	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type specificDate struct {
	Date types.Time
}

func (d *specificDate) Query() map[string]string {
	return map[string]string{
		"specificDate": d.Date.String(),
	}
}

func (d *specificDate) Headers() map[string]string {
	return map[string]string{}
}

func (d *specificDate) Body() map[string]net.BodyData {
	return map[string]net.BodyData{}
}

// TODO: determine unsuccessful response type
func (c *ClientImpl) TransactionHistoryDaily(date *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error) {
	payload := new(specificDate)

	if date != nil {
		payload.Date = types.NewTime(*date)
	}

	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/transactions/history/daily", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	txs, err := parseCsv(data)
	if err != nil {
		return nil, nil, err
	}

	return txs, nil, nil
}
