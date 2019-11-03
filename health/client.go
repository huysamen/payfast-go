package health

import "github.com/huysamen/payfast-go/types"

const pingPath = "/ping"

type Client struct {
	get types.GetRequest
}

func Create(get types.GetRequest) *Client {
	return &Client{get: get}
}
