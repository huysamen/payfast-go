package types

type CreditCardStatus struct {
	PayfastPaymentID  int    `json:"pf_payment_id"`
	MerchantPaymentID string `json:"m_payment_id"`
	Status            string `json:"status"`
	TransactionToken  string `json:"transaction_token"`
	Amount            int    `json:"amount"`
	CreditCardStatus  string `json:"cc_status"`
	CreditCardMessage string `json:"cc_message"`
}

func (c *CreditCardStatus) Copy(data map[string]interface{}) {
	c.PayfastPaymentID = int(data["pf_payment_id"].(float64))
	c.MerchantPaymentID = data["m_payment_id"].(string)
	c.Status = data["status"].(string)
	c.TransactionToken = data["transaction_token"].(string)
	c.Amount = int(data["amount"].(float64))
	c.CreditCardStatus = data["cc_status"].(string)
	c.CreditCardMessage = data["cc_message"].(string)
}
