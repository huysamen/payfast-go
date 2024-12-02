package net

type BodyData interface {
	Val() any
	String() string
}

type Payload interface {
	Headers() map[string]string
	Query() map[string]string
	Body() map[string]BodyData
}
