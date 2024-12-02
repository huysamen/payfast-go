package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type SubscriptionAdHocChargeRequest struct {
	Amount          types.Numeric
	ItemName        types.AlphaNumeric
	ItemDescription types.AlphaNumeric
	ITN             types.Bool
	MPaymentID      types.AlphaNumeric
	CreditCardCVV   types.Numeric
}

type SubscriptionAdHocChargeResponse struct {
	Response    bool   `json:"response"`
	Message     string `json:"message"`
	PFPaymentID string `json:"pf_payment_id"`
}

func (r *SubscriptionAdHocChargeRequest) Headers() map[string]string {
	return map[string]string{}
}

func (r *SubscriptionAdHocChargeRequest) Query() map[string]string {
	return map[string]string{}
}

func (r *SubscriptionAdHocChargeRequest) Body() map[string]net.BodyData {
	return map[string]net.BodyData{
		"amount":           r.Amount,
		"item_name":        r.ItemName,
		"item_description": r.ItemDescription,
		"itn":              r.ITN,
		"m_payment_id":     r.MPaymentID,
		"cc_cvv":           r.CreditCardCVV,
	}
}

func (c *ClientImpl) SubscriptionAdHocCharge(
	token string,
	payload *SubscriptionAdHocChargeRequest,
) (
	rsp *types.Response[SubscriptionAdHocChargeResponse],
	errRsp *types.ErrorResponse[int],
	err error,
) {
	data, _, err := net.Post(c.client, c.merchantID, c.merchantPassphrase, "/subscriptions/"+token+"/adhoc", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.Response[SubscriptionAdHocChargeResponse], types.ErrorResponse[int]](data, 200)
}
