package payfast

import (
	"bytes"
	"encoding/csv"
	"io"
	"time"

	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type dateRange struct {
	From types.Time
	To   types.Time
}

func (r *dateRange) Query() map[string]string {
	return map[string]string{
		"from": r.From.String(),
		"to":   r.To.String(),
	}
}

func (r *dateRange) Headers() map[string]string {
	return map[string]string{}
}

func (r *dateRange) Body() map[string]net.BodyData {
	return map[string]net.BodyData{}
}

// TODO: determine unsuccessful response type
func (c *ClientImpl) TransactionHistory(from *time.Time, to *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error) {
	payload := new(dateRange)

	if from != nil {
		payload.From = types.NewTime(*from)
	}

	if to != nil {
		payload.To = types.NewTime(*to)
	}

	data, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/transactions/history", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	txs, err := parseCsv(data)
	if err != nil {
		return nil, nil, err
	}

	return txs, nil, nil
}

func parseCsv(body []byte) ([]*types.Transaction, error) {
	reader := csv.NewReader(bytes.NewReader(body))
	txs := make([]*types.Transaction, 0)
	withBatch := false

	reader.LazyQuotes = true

	for {
		tx, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if tx[0] == "Date" {
			withBatch = tx[8] == "Batch ID"
			continue
		}

		t := new(types.Transaction)

		err = t.Copy(tx, withBatch)
		if err != nil {
			return nil, err
		}

		txs = append(txs, t)
	}

	return txs, nil
}
