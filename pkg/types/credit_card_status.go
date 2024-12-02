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
