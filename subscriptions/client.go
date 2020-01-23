package subscriptions

import (
	"strings"

	"github.com/huysamen/payfast-go/types"
)

const (
	basePath        = "/subscriptions/"
	fetchPath       = "/fetch"
	pausePath       = "/pause"
	unpausePath     = "/unpause"
	cancelPath      = "/cancel"
	updatePath      = "/update"
	adHocChargePath = "/adhoc"
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

func PathCat(base string, token string, action string) string {
	var b strings.Builder

	b.WriteString(base)
	b.WriteString(token)
	b.WriteString(action)

	return b.String()
}
