package net

import (
	"io"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.payfast.co.za"

func Get(client *http.Client, merchantID uint64, merchantPassphrase, url string, payload Payload, testing bool) ([]byte, int, error) {
	return do(client, merchantID, merchantPassphrase, http.MethodGet, url, payload, testing)
}

func Post(client *http.Client, merchantID uint64, merchantPassphrase, url string, payload Payload, testing bool) ([]byte, int, error) {
	return do(client, merchantID, merchantPassphrase, http.MethodPost, url, payload, testing)
}

func Put(client *http.Client, merchantID uint64, merchantPassphrase, url string, payload Payload, testing bool) ([]byte, int, error) {
	return do(client, merchantID, merchantPassphrase, http.MethodPut, url, payload, testing)
}

func PostURLEncoded(client *http.Client, merchantID uint64, merchantPassphrase, url string, payload *url.Values, testing bool) ([]byte, int, error) {
	return do(client, merchantID, merchantPassphrase, http.MethodPatch, url, payload, testing)
}

func do(client *http.Client, merchantID uint64, merchantPassphrase, method, url string, payload Payload, testing bool) ([]byte, int, error) {
	req, err := generateSignedRequest(merchantID, merchantPassphrase, method, baseUrl+url, payload, testing)
	if err != nil {
		return nil, 0, err
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer func() { _ = rsp.Body.Close() }()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, rsp.StatusCode, nil
}

func doURLEncoded(client *http.Client, merchantID uint64, merchantPassphrase, method, url string, payload *url.Values, testing bool) ([]byte, int, error) {
	req, err := generateSignedRequest(merchantID, merchantPassphrase, method, baseUrl+url, payload, testing)
	if err != nil {
		return nil, 0, err
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer func() { _ = rsp.Body.Close() }()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, rsp.StatusCode, nil
}
