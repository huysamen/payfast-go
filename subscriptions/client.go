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
	get   types.RemoteCall
	put   types.RemoteCall
	patch types.RemoteCall
	post  types.RemoteCall
}

func Create(get types.RemoteCall, put types.RemoteCall, patch types.RemoteCall, post types.RemoteCall) *Client {
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
