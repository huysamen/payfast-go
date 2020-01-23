package types

import "net/url"

type GetRequest func(url string, qp *url.Values) ([]byte, error)
type PutRequest func(url string, payload interface{}) ([]byte, error)
type PostRequest func(url string, payload interface{}) ([]byte, error)
type PatchRequest func(url string, payload interface{}) ([]byte, error)
