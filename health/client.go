package health

import (
	"github.com/huysamen/payfast-go/types"
)

const pingPath = "/ping"

type Client struct {
	get types.RemoteCall
}

func Create(get types.RemoteCall) *Client {
	return &Client{get: get}
}
