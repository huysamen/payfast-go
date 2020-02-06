package transactions

import (
	"github.com/huysamen/payfast-go/types"
)

const (
	historyPath = "/transactions/history"
	dailyPath   = "/transactions/history/daily"
	weeklyPath  = "/transactions/history/weekly"
	monthlyPath = "/transactions/history/monthly"
)

type Client struct {
	get types.RemoteCall
}

func Create(get types.RemoteCall) *Client {
	return &Client{get: get}
}
