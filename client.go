package payfast_go

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/huysamen/payfast-go/codec"
	"github.com/huysamen/payfast-go/health"
	"github.com/huysamen/payfast-go/subscriptions"
	"github.com/huysamen/payfast-go/transactions"
)

const baseUrl = "https://api.payfast.co.za"

type Client struct {
	merchantID         uint64
	merchantPassphrase string
	testing            bool
	http               *http.Client
	Health             *health.Client
	Subscriptions      *subscriptions.Client
	Transactions       *transactions.Client
}

func New(merchantID uint64, merchantPassphrase string, client *http.Client, testing bool) *Client {
	if client == nil {
		client = defaultClient()
	}

	pf := &Client{
		merchantID:         merchantID,
		merchantPassphrase: merchantPassphrase,
		testing:            testing,
		http:               client,
	}

	pf.createServices()

	return pf
}

func Default(testing bool) (*Client, error) {
	ID := os.Getenv("PAYFAST_MERCHANT_ID")
	if ID == "" {
		return nil, errors.New("no api merchant ID present")
	}

	IDi, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, err
	}

	passphrase := os.Getenv("PAYFAST_MERCHANT_PASSPHRASE")
	if passphrase == "" {
		return nil, errors.New("no api merchant passphrase present")
	}

	return New(IDi, passphrase, defaultClient(), testing), nil
}

func defaultClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
}

func (a *Client) createServices() {
	a.Health = health.Create(a.get)
	a.Subscriptions = subscriptions.Create(a.get, a.put, a.patch, a.post)
	a.Transactions = transactions.Create(a.get)
}

func (a *Client) get(path string, payload any) ([]byte, int, error) {
	req, err := codec.GenerateSignedRequest(a.merchantID, a.merchantPassphrase, "GET", baseUrl+path, payload, a.testing)
	if err != nil {
		return nil, 0, err
	}

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer func() { _ = rsp.Body.Close() }()

	// todo: check auth here

	body, err := io.ReadAll(rsp.Body)

	return body, rsp.StatusCode, err
}

func (a *Client) put(path string, payload any) ([]byte, int, error) {
	return a.putPostPatch("PUT", path, payload)
}

func (a *Client) post(path string, payload any) ([]byte, int, error) {
	return a.putPostPatch("POST", path, payload)
}

func (a *Client) patch(path string, payload any) ([]byte, int, error) {
	return a.putPostPatch("PATCH", path, payload)
}

func (a *Client) putPostPatch(method string, path string, data any) ([]byte, int, error) {
	req, err := codec.GenerateSignedRequest(a.merchantID, a.merchantPassphrase, method, baseUrl+path, data, a.testing)
	if err != nil {
		return nil, 0, err
	}

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer func() { _ = rsp.Body.Close() }()

	// todo: check auth here

	body, err := io.ReadAll(rsp.Body)

	return body, rsp.StatusCode, err
}
