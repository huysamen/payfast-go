package types

type Response struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Data   ResponseData `json:"data"`
}

type ResponseData struct {
	Response any `json:"response"`
}
