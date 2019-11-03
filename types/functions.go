package types

type GetRequest func(url string) ([]byte, error)
type PostRequest func(url string, payload interface{}) ([]byte, error)
