package types

import "github.com/huysamen/payfast-go/utils/copyutils"

type CreditCardStatus struct {
	PayfastPaymentID  int    `json:"pf_payment_id"`
	MerchantPaymentID string `json:"m_payment_id"`
	Status            string `json:"status"`
	TransactionToken  string `json:"transaction_token"`
	Amount            int    `json:"amount"`
	CreditCardStatus  string `json:"cc_status"`
	CreditCardMessage string `json:"cc_message"`
}

func (c *CreditCardStatus) Copy(data map[string]any) {
	c.PayfastPaymentID = copyutils.CopyInt(data, "pf_payment_id")
	c.MerchantPaymentID = copyutils.CopyString(data, "m_payment_id")
	c.Status = copyutils.CopyString(data, "status")
	c.TransactionToken = copyutils.CopyString(data, "transaction_token")
	c.Amount = copyutils.CopyInt(data, "amount")
	c.CreditCardStatus = copyutils.CopyString(data, "cc_status")
	c.CreditCardMessage = copyutils.CopyString(data, "cc_message")
}
