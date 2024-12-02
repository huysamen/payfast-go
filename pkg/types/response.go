package types

type Response[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Response T      `json:"response"`
		Message  string `json:"message"`
	} `json:"data"`
}

type ConfirmationResponse Response[bool]
type ErrorResponse[T any] Response[ErrorResponseData[T]]

type ErrorResponseData[T any] struct {
	Response T      `json:"response"`
	Message  string `json:"message"`
}
