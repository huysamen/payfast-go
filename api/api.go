package api

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/huysamen/payfast-go/health"
	"github.com/huysamen/payfast-go/subscriptions"
	"github.com/huysamen/payfast-go/types"
	"github.com/huysamen/payfast-go/utils/timeutils"
)

const baseUrl = "https://api.payfast.co.za"

type Api struct {
	merchantID         uint64
	merchantPassphrase string
	testing            bool
	http               *http.Client

	Health        *health.Client
	Subscriptions *subscriptions.Client
}

func Default() (*Api, error) {
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
		testing:            false,
		http: &http.Client{
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
	}

	api.createServices()

	return api, nil
}

func (a *Api) createServices() {
	a.Health = health.Create(a.get)
	a.Subscriptions = subscriptions.Create(a.get, a.put, a.patch, a.post)
}

func (a *Api) get(path string) ([]byte, error) {
	headers, _, err := a.generatePayload(nil)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", baseUrl+path, nil)
	req.Header = headers

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

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

func (a *Api) putPostPatch(method string, path string, payload interface{}) ([]byte, error) {
	headers, body, err := a.generatePayload(payload)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, baseUrl+path, bytes.NewBuffer(body))
	req.Header = headers
	req.Header.Set("content-type", "application/json")

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}

// todo: very unoptimised, just trying to match Payfast's non-standard signature creation
func (a *Api) generatePayload(data interface{}) (headers http.Header, payload []byte, err error) {
	fields := []string{"merchant-id", "version", "timestamp", "passphrase"}
	values := make(map[string]string)
	headers = make(http.Header)
	body := make(map[string]interface{})

	values["merchant-id"] = strconv.FormatUint(a.merchantID, 10)
	values["version"] = "v1"
	values["timestamp"] = timeutils.ToStandardString(time.Now())
	values["passphrase"] = a.merchantPassphrase

	headers["merchant-id"] = []string{strconv.FormatUint(a.merchantID, 10)}
	headers["version"] = []string{"v1"}
	headers["timestamp"] = []string{timeutils.ToStandardString(time.Now())}
	headers["passphrase"] = []string{a.merchantPassphrase}

	if data != nil {
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)
			tag := field.Tag.Get("payfast")
			attr := strings.Split(tag, ",")

			if len(attr) != 4 {
				return nil, nil, errors.New("incorrect payfast attributes format")
			}

			rv, sv, err := parseValues(attr[0], attr[2], attr[3], value.Interface())
			if err != nil {
				return nil, nil, err
			}

			switch attr[1] {
			case "header":
				if sv.Valid && sv.Value != "" {
					headers[attr[0]] = []string{sv.Value}
				}
			case "body":
				if rv != nil {
					body[attr[0]] = rv
				}
			}

			if sv.Valid && sv.Value != "" {
				fields = append(fields, attr[0])
				values[attr[0]] = sv.Value
			}
		}

		payload, err = json.Marshal(body)
		if err != nil {
			return nil, nil, err
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

	return headers, payload, nil
}

func parseValues(field string, fieldType string, required string, value interface{}) (interface{}, types.AlphaNumeric, error) {
	switch fieldType {
	case "numeric":
		n := value.(types.Numeric)

		if !n.Valid && required == "required" {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid {
			return n.Value, types.NewAlphaNumeric(strconv.Itoa(n.Value)), nil
		}

	case "yyyy-mm-dd":
		n := value.(types.AlphaNumeric)

		if (!n.Valid || n.Value == "") && required == "required" {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid && n.Value != "" {
			_, err := time.Parse("2006-01-02", n.Value)
			if err != nil {
				return nil, types.AlphaNumeric{}, err
			}

			return n.Value, n, nil
		}

	case "alphanumeric":
		n := value.(types.AlphaNumeric)

		if (!n.Valid || n.Value == "") && required == "required" {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid && n.Value != "" {
			return n.Value, n, nil
		}

	case "bool":
		n := value.(types.Bool)

		if !n.Valid && required == "required" {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid && n.Value {
			return n.Value, types.NewAlphaNumeric("1"), nil
		} else if n.Valid && !n.Value {
			return n.Value, types.NewAlphaNumeric("0"), nil
		}
	}

	return nil, types.AlphaNumeric{}, nil
}
