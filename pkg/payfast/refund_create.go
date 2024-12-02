package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type RefundCreateRequest struct {
	Amount            types.Numeric      `json:"amount"`
	Reason            types.AlphaNumeric `json:"reason"`
	NotifyMerchant    types.Bool         `json:"notify_merchant"`
	NotifyBuyer       types.Bool         `json:"notify_buyer"`
	BankAccountHolder types.AlphaNumeric `json:"bank_account_holder"`
	BankName          types.AlphaNumeric `json:"bank_name"`
	BankBranchCode    types.Numeric      `json:"bank_branch_code"`
	BankAccountNumber types.Numeric      `json:"bank_account_number"`
	BankAccountType   types.AlphaNumeric `json:"bank_account_type"` // current | savings
}

func (r *RefundCreateRequest) Headers() map[string]string {
	return map[string]string{}
}

func (r *RefundCreateRequest) Query() map[string]string {
	return map[string]string{}
}

func (r *RefundCreateRequest) Body() map[string]net.BodyData {
	return map[string]net.BodyData{
		"amount":              r.Amount,
		"reason":              r.Reason,
		"notify_merchant":     r.NotifyMerchant,
		"notify_buyer":        r.NotifyBuyer,
		"bank_account_holder": r.BankAccountHolder,
		"bank_name":           r.BankName,
		"bank_branch_code":    r.BankBranchCode,
		"bank_account_number": r.BankAccountNumber,
		"bank_account_type":   r.BankAccountType,
	}
}

// TODO: determine unsuccessful response type
func (c *ClientImpl) RefundCreate(pfPaymentID string, payload *RefundCreateRequest) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Post(c.client, c.merchantID, c.merchantPassphrase, "/refunds/"+pfPaymentID, payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.ConfirmationResponse, types.ErrorResponse[int]](data, 200)
}
