package net

import "encoding/json"

type BaseResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func ParseResponse[R any, E any](data []byte, success int) (*R, *E, error) {
	g := new(BaseResponse)

	err := json.Unmarshal(data, g)
	if err != nil {
		return nil, nil, err
	}

	if g.Code != success {
		errRsp := new(E)

		err = json.Unmarshal(data, errRsp)
		if err != nil {
			return nil, nil, err
		}

		return nil, errRsp, nil
	}

	rsp := new(R)

	err = json.Unmarshal(data, rsp)
	if err != nil {
		return nil, nil, err
	}

	return rsp, nil, nil
}
