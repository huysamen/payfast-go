package transactions

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/url"
	"time"

	"github.com/huysamen/payfast-go/types"
	"github.com/huysamen/payfast-go/utils/timeutils"
)

func (c *Client) History(from *time.Time, to *time.Time) ([]*types.Transaction, error) {
	qp := url.Values{}

	if from != nil {
		qp.Add("from", timeutils.ToDayString(*from))
	}

	if to != nil {
		qp.Add("to", timeutils.ToDayString(*to))
	}

	body, err := c.get(historyPath, &qp)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(body))
	txs := make([]*types.Transaction, 0)

	for {
		tx, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if tx[0] == "Date" {
			continue
		}

		t := new(types.Transaction)

		err = t.Copy(tx)
		if err != nil {
			return nil, err
		}

		txs = append(txs, t)
	}

	return txs, nil
}
