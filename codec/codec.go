package codec

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/huysamen/payfast-go/types"
	"github.com/huysamen/payfast-go/utils/timeutils"
)

func GenerateSignedRequest(
	merchantID uint64,
	merchantPassphrase string,
	method string,
	path string,
	data interface{},
	testing bool,
) (*http.Request, error) {
	query, headers, body, err := sign(merchantID, merchantPassphrase, data)
	if err != nil {
		return nil, err
	}

	(*headers)["content-type"] = []string{"application/json"}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(*body))
	req.Header = *headers

	if testing == true {
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
	data interface{},
) (*url.Values, *http.Header, *[]byte, error) {
	fields := []string{"merchant-id", "version", "timestamp", "passphrase"}
	values := make(map[string]string)

	query := make(url.Values)
	headers := make(http.Header)
	body := make(map[string]interface{})

	values["merchant-id"] = strconv.FormatUint(merchantID, 10)
	values["version"] = "v1"
	values["timestamp"] = timeutils.ToStandardString(time.Now())
	values["passphrase"] = merchantPassphrase

	headers["merchant-id"] = []string{strconv.FormatUint(merchantID, 10)}
	headers["version"] = []string{"v1"}
	headers["timestamp"] = []string{timeutils.ToStandardString(time.Now())}
	headers["passphrase"] = []string{merchantPassphrase}

	var payload []byte

	if data != nil {
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)
			tag := field.Tag.Get("payfast")
			attr := strings.Split(tag, ",")

			if len(attr) != 4 {
				return nil, nil, nil, errors.New("incorrect payfast attributes format")
			}

			rv, sv, err := parseValues(attr[0], attr[2], attr[3], value.Interface())
			if err != nil {
				return nil, nil, nil, err
			}

			switch attr[1] {

			case "query":
				if sv.Valid && sv.Value != "" {
					query.Add(attr[0], sv.Value)
				}

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

		var err error

		payload, err = json.Marshal(body)
		if err != nil {
			return nil, nil, nil, err
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

	return &query, &headers, &payload, nil
}

func parseValues(field string, fieldType string, requiredStr string, value interface{}) (interface{}, types.AlphaNumeric, error) {
	required := requiredStr == "required"

	switch fieldType {
	case "numeric":
		n := value.(types.Numeric)

		if !n.Valid && required {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid {
			return n.Value, types.NewAlphaNumeric(strconv.Itoa(n.Value)), nil
		}

	case "yyyy-mm-dd":
		n := value.(types.Time)

		if !n.Valid && required {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid {
			return n.Value, types.NewAlphaNumeric(n.Value.Format("2006-01-02")), nil
		}

	case "alphanumeric":
		n := value.(types.AlphaNumeric)

		if (!n.Valid || n.Value == "") && required {
			return nil, types.AlphaNumeric{}, errors.New("field " + field + " is required")
		}

		if n.Valid && n.Value != "" {
			return n.Value, n, nil
		}

	case "bool":
		n := value.(types.Bool)

		if !n.Valid && required {
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
