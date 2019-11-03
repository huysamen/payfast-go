package subscriptions

import "github.com/huysamen/payfast-go/types"

const (
	fetchPath = "/subscriptions/__sid__/fetch"
)

type Client struct {
	get types.GetRequest
}

func Create(get types.GetRequest) *Client {
	return &Client{get: get}
}
