package transactions

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/huysamen/payfast-go/types"
)

func parseCsv(body []byte) ([]*types.Transaction, error) {
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
