package types

type RemoteCall func(url string, payload any) (rsp []byte, status int, err error)
