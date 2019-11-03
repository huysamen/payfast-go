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
	a.Subscriptions = subscriptions.Create(a.get)
}

func (a *Api) get(path string) ([]byte, error) {
	headers, _ := a.generatePayload(nil)
	req, _ := http.NewRequest("GET", baseUrl+path, nil)

	req.Header = headers

	rsp, err := a.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}

func (a *Api) post(path string, payload interface{}) ([]byte, error) {
	pt := reflect.ValueOf(payload).Elem()
	f := pt.FieldByName("Key")

	if f.IsValid() {
		f.SetString(a.merchantPassphrase)
	}

	req, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	rsp, err := a.http.Post(baseUrl+path, "application/json", bytes.NewBuffer(req))

	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}

// todo: very unoptimised, just trying to match Payfast's non-standard signature creation
func (a *Api) generatePayload(data interface{}) (headers http.Header, body map[string]interface{}) {
	fields := []string{"merchant-id", "version", "timestamp", "passphrase"}
	values := make(map[string]string)
	headers = make(http.Header)
	body = make(map[string]interface{})

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

			if strings.Contains(tag, "header") {
				headers[field.Name] = []string{value.String()}
			} else if strings.Contains(tag, "body") {
				body[field.Name] = value.Interface()
			}

			fields = append(fields, field.Name)
			values[field.Name] = value.String()
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

	return headers, body
}
