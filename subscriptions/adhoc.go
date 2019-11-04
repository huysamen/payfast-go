package subscriptions

import (
	"encoding/json"
	"strings"

	"github.com/huysamen/payfast-go/types"
)

type AdHocSubscriptionChargeReq struct {
	Amount          types.Numeric      `payfast:"amount,body,numeric,required"`                // The amount which the buyer must pay, in CENTS (ZAR).
	ItemName        types.AlphaNumeric `payfast:"item_name,body,alphanumeric,required"`        // The name of the item being charged for.
	ItemDescription types.AlphaNumeric `payfast:"item_description,body,alphanumeric,optional"` // The name of the item being charged for.
	ITN             types.Bool         `payfast:"itn,body,bool,optional"`                      // Specify whether an ITN must be sent for the ad hoc charge (1 by default).
	MPaymentID      types.AlphaNumeric `payfast:"m_payment_id,body,alphanumeric,optional"`     // Unique payment ID on the merchant’s system.
	CreditCardCVV   types.Numeric      `payfast:"cc_cvv,body,numeric,optional"`                // The credit card cvv number.
}

func (c *Client) AdHocCharge(subscriptionID string, payload AdHocSubscriptionChargeReq) (bool, error) {
	body, err := c.post(strings.ReplaceAll(adHocChargePath, "__sid__", subscriptionID), payload)
	if err != nil {
		return false, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return false, err
	}

	if rsp.Code == 200 {
		return true, nil
	}

	return false, nil
}
