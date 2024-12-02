package payfast

import "github.com/huysamen/payfast-go/internal/net"

func (c *ClientImpl) Ping() (up bool, err error) {
	rsp, _, err := net.Get(c.client, c.merchantID, c.merchantPassphrase, "/ping", nil, c.testing)
	if err != nil {
		return false, err
	}

	return string(rsp) == "\"PayFast API\"", nil
}
