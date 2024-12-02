package payfast

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/huysamen/payfast-go/pkg/types"
)

type Client interface {
	Ping() (up bool, err error)

	SubscriptionFetch(token string) (rsp *types.Response[types.Subscription], errRsp *types.ErrorResponse[int], err error)
	SubscriptionPause(token string) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error)
	SubscriptionUnpause(token string) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error)
	SubscriptionCancel(token string) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error)
	SubscriptionUpdate(token string, payload *SubscriptionUpdateRequest) (rsp *types.Response[types.Subscription], errRsp *types.ErrorResponse[int], err error)
	SubscriptionAdHocCharge(token string, payload *SubscriptionAdHocChargeRequest) (rsp *types.Response[SubscriptionAdHocChargeResponse], errRsp *types.ErrorResponse[int], err error)

	TransactionHistory(from *time.Time, to *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error)
	TransactionHistoryDaily(date *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error)
	TransactionHistoryWeekly(date *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error)
	TransactionHistoryMonthly(date *time.Time) (rsp []*types.Transaction, errRsp *types.ErrorResponse[int], err error)

	CardTransactionQuery(idOrToken string) (rsp *types.Response[types.CreditCardStatus], errRsp *types.ErrorResponse[int], err error)

	RefundQuery(pfPaymentID string) (rsp *types.RefundQuery, errRsp *types.ErrorResponse[int], err error)
	RefundCreate(pfPaymentID string, payload *RefundCreateRequest) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error)
	RefundRetrieve(pfPaymentID string) (rsp *types.Response[types.Refund], errRsp *types.ErrorResponse[int], err error)
}

type ClientImpl struct {
	merchantID         uint64
	merchantPassphrase string
	testing            bool
	client             *http.Client
}

func Default(testing bool) (Client, error) {
	envID := os.Getenv("PAYFAST_MERCHANT_ID")
	if envID == "" {
		return nil, errors.New("no api merchant envID present")
	}

	id, err := strconv.ParseUint(envID, 10, 64)
	if err != nil {
		return nil, err
	}

	passphrase := os.Getenv("PAYFAST_MERCHANT_PASSPHRASE")
	if passphrase == "" {
		return nil, errors.New("no api merchant passphrase present")
	}

	return NewWithClient(id, passphrase, testing, defaultClient()), nil
}

func New(merchantID uint64, merchantPassphrase string, testing bool) Client {
	return NewWithClient(merchantID, merchantPassphrase, testing, defaultClient())
}

func NewWithClient(merchantID uint64, merchantPassphrase string, testing bool, client *http.Client) Client {
	pf := &ClientImpl{
		merchantID:         merchantID,
		merchantPassphrase: merchantPassphrase,
		testing:            testing,
		client:             client,
	}

	return pf
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
