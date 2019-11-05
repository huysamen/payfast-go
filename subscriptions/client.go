package subscriptions

import "github.com/huysamen/payfast-go/types"

const (
	fetchPath       = "/subscriptions/__token__/fetch"
	pausePath       = "/subscriptions/__token__/pause"
	unpausePath     = "/subscriptions/__token__/unpause"
	cancelPath      = "/subscriptions/__token__/cancel"
	updatePath      = "/subscriptions/__token__/update"
	adHocChargePath = "/subscriptions/__token__/adhoc"
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
