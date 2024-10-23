package types

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/huysamen/payfast-go/utils/timeutils"
)

type Transaction struct {
	Date          time.Time
	Type          string
	Sign          string
	Party         string
	Name          string
	Description   string
	Currency      string
	FundingType   string
	Gross         float64
	Fee           float64
	Net           float64
	Balance       float64
	MPaymentID    string
	PFPaymentID   string
	CustomString1 string
	CustomInt1    int
	CustomString2 string
	CustomInt2    int
	CustomString3 string
	CustomString4 string
	CustomString5 string
	CustomInt3    int
	CustomInt4    int
	CustomInt5    int
	BatchID       string
}

func (t *Transaction) Copy(tx []string, withBatch bool) error {
	offset := 0

	if withBatch {
		offset = 1
	}

	if len(tx) != 24+offset {
		return errors.New("incorrect csv data: " + strings.Join(tx, ","))
	}

	t.Date = timeutils.FromCsvString(tx[0])
	t.Type = tx[1]
	t.Sign = tx[2]
	t.Party = tx[3]
	t.Name = tx[4]
	t.Description = tx[5]
	t.Currency = tx[6]
	t.FundingType = tx[7]

	if withBatch {
		t.BatchID = tx[8]
	}

	if tx[8+offset] != "" {
		f64, err := copyFloat64(tx, 8+offset)
		if err != nil {
			return err
		}

		t.Gross = f64
	}

	if tx[9+offset] != "" {
		f64, err := copyFloat64(tx, 9+offset)
		if err != nil {
			return err
		}

		t.Fee = f64
	}

	if tx[10+offset] != "" {
		f64, err := copyFloat64(tx, 10+offset)
		if err != nil {
			return err
		}

		t.Net = f64
	}

	if tx[11+offset] != "" {
		f64, err := copyFloat64(tx, 11+offset)
		if err != nil {
			return err
		}

		t.Balance = f64
	}

	t.MPaymentID = tx[12+offset]
	t.PFPaymentID = tx[13+offset]
	t.CustomString1 = tx[14+offset]
	t.CustomString2 = tx[16+offset]
	t.CustomString3 = tx[18+offset]
	t.CustomString4 = tx[19+offset]
	t.CustomString5 = tx[20+offset]

	if tx[15+offset] != "" {
		i, err := strconv.Atoi(tx[15+offset])
		if err != nil {
			return err
		}

		t.CustomInt1 = i
	}

	if tx[17+offset] != "" {
		i, err := strconv.Atoi(tx[17+offset])
		if err != nil {
			return err
		}

		t.CustomInt2 = i
	}

	if tx[21+offset] != "" {
		i, err := strconv.Atoi(tx[21+offset])
		if err != nil {
			return err
		}

		t.CustomInt3 = i
	}

	if tx[22+offset] != "" {
		i, err := strconv.Atoi(tx[22+offset])
		if err != nil {
			return err
		}

		t.CustomInt4 = i
	}

	if tx[23+offset] != "" {
		i, err := strconv.Atoi(tx[23+offset])
		if err != nil {
			return err
		}

		t.CustomInt5 = i
	}

	return nil
}

func copyFloat64(data []string, idx int) (float64, error) {
	if data[idx] == "" {
		return 0.0, nil
	}

	return strconv.ParseFloat(strings.ReplaceAll(data[idx], ",", ""), 64)
}
