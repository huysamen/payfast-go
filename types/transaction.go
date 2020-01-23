package types

import (
	"github.com/huysamen/payfast-go/utils/timeutils"
	"strconv"
	"strings"
	"time"
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
	CustomString2 string
	CustomString3 string
	CustomString4 string
	CustomString5 string
	CustomInt1    int
	CustomInt2    int
	CustomInt3    int
	CustomInt4    int
	CustomInt5    int
}

func (t *Transaction) Copy(tx []string) error {
	t.Date = timeutils.FromCsvString(tx[0])
	t.Type = tx[1]
	t.Sign = tx[2]
	t.Party = tx[3]
	t.Name = tx[4]
	t.Description = tx[5]
	t.Currency = tx[6]
	t.FundingType = tx[7]

	if tx[8] != "" {
		f64, err := strconv.ParseFloat(strings.ReplaceAll(tx[8], ",", ""), 64)
		if err != nil {
			return err
		}
		t.Gross = f64
	}

	if tx[9] != "" {
		f64, err := strconv.ParseFloat(strings.ReplaceAll(tx[9], ",", ""), 64)
		if err != nil {
			return err
		}
		t.Fee = f64
	}

	if tx[10] != "" {
		f64, err := strconv.ParseFloat(strings.ReplaceAll(tx[10], ",", ""), 64)
		if err != nil {
			return err
		}
		t.Net = f64
	}

	if tx[11] != "" {
		f64, err := strconv.ParseFloat(strings.ReplaceAll(tx[11], ",", ""), 64)
		if err != nil {
			return err
		}
		t.Balance = f64
	}

	t.MPaymentID = tx[12]
	t.PFPaymentID = tx[13]
	t.CustomString1 = tx[14]
	t.CustomString2 = tx[15]
	t.CustomString3 = tx[16]
	t.CustomString4 = tx[17]
	t.CustomString5 = tx[18]

	if tx[19] != "" {
		i, err := strconv.Atoi(tx[19])
		if err != nil {
			return err
		}
		t.CustomInt1 = i
	}

	if tx[19] != "" {
		i, err := strconv.Atoi(tx[20])
		if err != nil {
			return err
		}
		t.CustomInt2 = i
	}

	if tx[19] != "" {
		i, err := strconv.Atoi(tx[21])
		if err != nil {
			return err
		}
		t.CustomInt3 = i
	}

	if tx[19] != "" {
		i, err := strconv.Atoi(tx[22])
		if err != nil {
			return err
		}
		t.CustomInt4 = i
	}

	if tx[19] != "" {
		i, err := strconv.Atoi(tx[23])
		if err != nil {
			return err
		}
		t.CustomInt5 = i
	}

	return nil
}
