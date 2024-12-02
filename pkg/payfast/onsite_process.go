package payfast

import (
	"github.com/huysamen/payfast-go/internal/net"
	"github.com/huysamen/payfast-go/pkg/types"
)

type OnsiteProcessRequest struct {
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

func (r *OnsiteProcessRequest) Headers() map[string]string {
	return map[string]string{}
}

func (r *OnsiteProcessRequest) Query() map[string]string {
	return map[string]string{}
}

func (r *OnsiteProcessRequest) Body() map[string]net.BodyData {
	return map[string]net.BodyData{}
}

// TODO: determine unsuccessful response type
func (c *ClientImpl) OnsiteProcess(payload *OnsiteProcessRequest) (rsp *types.ConfirmationResponse, errRsp *types.ErrorResponse[int], err error) {
	data, _, err := net.Post(c.client, c.merchantID, c.merchantPassphrase, "/onsite/process", payload, c.testing)
	if err != nil {
		return nil, nil, err
	}

	return net.ParseResponse[types.ConfirmationResponse, types.ErrorResponse[int]](data, 200)
}
