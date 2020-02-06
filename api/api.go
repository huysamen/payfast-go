package api

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
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

type Api struct {
	merchantID         uint64
	merchantPassphrase string
	testing            bool
	http               *http.Client
	Health             *health.Client
	Subscriptions      *subscriptions.Client
	Transactions       *transactions.Client
}

func New(client *http.Client, testing bool) (*Api, error) {
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

	api := &Api{
		merchantID:         IDi,
		merchantPassphrase: passphrase,
		testing:            testing,
		http:               client,
	}

	api.createServices()

	return api, nil
}

func Default(testing bool) (*Api, error) {
	return New(
		&http.Client{
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
		},
		testing,
	)
}

func (a *Api) createServices() {
	a.Health = health.Create(a.get)
	a.Subscriptions = subscriptions.Create(a.get, a.put, a.patch, a.post)
	a.Transactions = transactions.Create(a.get)
}

func (a *Api) get(path string, payload interface{}) ([]byte, error) {
	req, err := codec.GenerateSignedRequest(a.merchantID, a.merchantPassphrase, "GET", baseUrl+path, payload, a.testing)
	if err != nil {
		return nil, err
	}

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	// todo: check auth here

	return ioutil.ReadAll(rsp.Body)
}

func (a *Api) put(path string, payload interface{}) ([]byte, error) {
	return a.putPostPatch("PUT", path, payload)
}

func (a *Api) post(path string, payload interface{}) ([]byte, error) {
	return a.putPostPatch("POST", path, payload)
}

func (a *Api) patch(path string, payload interface{}) ([]byte, error) {
	return a.putPostPatch("PATCH", path, payload)
}

func (a *Api) putPostPatch(method string, path string, data interface{}) ([]byte, error) {
	req, err := codec.GenerateSignedRequest(a.merchantID, a.merchantPassphrase, method, baseUrl+path, data, a.testing)
	if err != nil {
		return nil, err
	}

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	// todo: check auth here

	return ioutil.ReadAll(rsp.Body)
}
