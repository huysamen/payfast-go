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
	get   types.GetRequest
	put   types.PutRequest
	patch types.PatchRequest
	post  types.PostRequest
}

func Create(get types.GetRequest) *Client {
	return &Client{get: get}
}
