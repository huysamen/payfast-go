package health

func (c *Client) Ping() (bool, error) {
	rsp, _, err := c.get(pingPath, nil)
	if err != nil {
		return false, err
	}

	return string(rsp) == "\"PayFast API\"", nil
}
