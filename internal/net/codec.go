package net

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/huysamen/payfast-go/internal/utils/timeutils"
)

func generateSignedRequest(
	merchantID uint64,
	merchantPassphrase string,
	method string,
	path string,
	payload Payload,
	testing bool,
) (*http.Request, error) {
	query, headers, body, err := sign(merchantID, merchantPassphrase, payload)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(*body))
	req.Header = *headers

	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}

	if testing {
		if query == nil {
			query = new(url.Values)
		}

		query.Add("testing", "true")
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	return req, nil
}

func sign(
	merchantID uint64,
	merchantPassphrase string,
	payload Payload,
) (*url.Values, *http.Header, *[]byte, error) {
	fields := []string{"merchant-id", "version", "timestamp", "passphrase"}

	values := make(map[string]string)
	values["merchant-id"] = strconv.FormatUint(merchantID, 10)
	values["version"] = "v1"
	values["timestamp"] = timeutils.ToStandardString(time.Now())
	values["passphrase"] = merchantPassphrase

	headers := make(http.Header)
	headers["merchant-id"] = []string{strconv.FormatUint(merchantID, 10)}
	headers["version"] = []string{"v1"}
	headers["timestamp"] = []string{timeutils.ToStandardString(time.Now())}
	headers["passphrase"] = []string{merchantPassphrase}

	query := make(url.Values)
	body := make(map[string]any)

	var data []byte

	if payload != nil {
		for k, v := range payload.Headers() {
			if v != "" {
				headers[k] = []string{v}
				fields = append(fields, k)
				values[k] = v
			}
		}

		for k, v := range payload.Query() {
			if v != "" {
				query.Add(k, v)
				fields = append(fields, k)
				values[k] = v
			}
		}

		for k, v := range payload.Body() {
			if v.Value() != nil && v.String() != "" {
				body[k] = v.Value()
				fields = append(fields, k)
				values[k] = v.String()
			}
		}

		if len(body) > 0 {
			var err error

			data, err = json.Marshal(body)
			if err != nil {
				return nil, nil, nil, err
			}
		}
	}

	sort.Strings(fields)

	signature := ""

	for _, f := range fields {
		v := values[f]

		if len(v) > 0 {
			signature = signature + "&" + url.QueryEscape(f) + "=" + url.QueryEscape(v)
		}
	}

	signature = signature[1:]

	hash := md5.New()
	_, _ = hash.Write([]byte(signature))
	sig := hex.EncodeToString(hash.Sum(nil))

	headers["signature"] = []string{sig}

	return &query, &headers, &data, nil
}

func signURLRequest(merchantPassphrase string, data map[string]string) (payload string, err error) {
	data["passphrase"] = merchantPassphrase
	keys := make([]string, len(data))

	i := 0

	for key := range data {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	var signature string

	for _, key := range keys {
		signature += "&" + url.QueryEscape(key) + "=" + url.QueryEscape(data[key])
	}

	signature = signature[1:]

	//nolint:gosec
	hash := md5.New()
	_, _ = hash.Write([]byte(signature))
	sig := hex.EncodeToString(hash.Sum(nil))

	p := url.Values{}

	for _, k := range keys {
		p.Add(k, data[k])
	}

	p.Add("signature", sig)

	return p.Encode(), nil
}
