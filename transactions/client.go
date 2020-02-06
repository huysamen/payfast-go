package transactions

import (
	"strings"

	"github.com/huysamen/payfast-go/types"
)

const (
	historyPath = "/transactions/history"
	dailyPath   = "/transactions/history/daily"
	weeklyPath  = "/transactions/history/weekly"
	monthlyPath = "/transactions/history/monthly"
	queryPath   = "/process/query/"
)

type Client struct {
	get types.RemoteCall
}

func Create(get types.RemoteCall) *Client {
	return &Client{get: get}
}

func PathCat(base string, token string) string {
	var b strings.Builder

	b.WriteString(base)
	b.WriteString(token)

	return b.String()
}
