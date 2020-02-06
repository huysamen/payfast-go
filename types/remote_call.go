package types

type RemoteCall func(url string, payload interface{}) ([]byte, error)
