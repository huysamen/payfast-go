package subscriptions

import "github.com/huysamen/payfast-go/types"

const (
	fetchPath       = "/subscriptions/__sid__/fetch"
	pausePath       = "/subscriptions/__sid__/pause"
	unpausePath     = "/subscriptions/__sid__/unpause"
	cancelPath      = "/subscriptions/__sid__/cancel"
	updatePath      = "/subscriptions/__sid__/update"
	adHocChargePath = "/subscriptions/__sid__/adhoc"
)

type Client struct {
	get   types.GetRequest
	put   types.PutRequest
	patch types.PatchRequest
	post  types.PostRequest
}

func Create(get types.GetRequest, put types.PutRequest, patch types.PatchRequest, post types.PostRequest) *Client {
	return &Client{
		get:   get,
		put:   put,
		patch: patch,
		post:  post,
	}
}
