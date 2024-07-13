package types

type RemoteCall func(url string, payload any) ([]byte, error)
