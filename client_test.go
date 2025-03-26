package payfast_go

import (
	"fmt"
	"testing"
	"time"

	"github.com/huysamen/payfast-go/transactions"
	"github.com/huysamen/payfast-go/types"
)

func TestClient_createServices(t *testing.T) {
	c := New(10227874, "vbai3j0ojfhqi", nil, false)

	txs, status, err := c.Transactions.History(
		transactions.TransactionHistoryReq{
			From: types.NewTime(time.Now().Add(-time.Hour)),
			To:   types.NewTime(time.Now()),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(txs, status)
}
