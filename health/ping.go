package health

func (c *Client) Ping() (bool, error) {
	rsp, err := c.get(pingPath)
	if err != nil {
		return false, err
	}

	return string(rsp) == "\"PayFast API\"", nil
}
