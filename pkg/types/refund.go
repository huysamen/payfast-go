package types

import "time"

type RefundQuery struct {
	// REFUNDABLE, COMPLETED, NOT_AVAILABLE
	Status                   string        `json:"status"`
	Token                    string        `json:"token"`
	FundingType              string        `json:"funding_type"`
	AmountOriginal           int           `json:"amount_original"`
	AmountAvailableForRefund int           `json:"amount_available_for_refund"`
	Errors                   []string      `json:"errors"`
	RefundFull               RefundFull    `json:"refund_full"`
	RefundPartial            RefundPartial `json:"refund_partial"`
	BankNames                BankNames     `json:"bank_names"`
}

type RefundFull struct {
	// PAYMENT_SOURCE, BANK_PAYOUT, NOT_AVAILABLE
	Method            string `json:"method"`
	BankAccountHolder string `json:"bank_account_holder"`
	BankName          string `json:"bank_name"`
	BankBranchCode    string `json:"bank_branch_code"`
	BankAccountNumber string `json:"bank_account_number"`
	BankAccountType   string `json:"bank_account_type"`
}

type RefundPartial struct {
	// PAYMENT_SOURCE, BANK_PAYOUT, NOT_AVAILABLE
	Method            string `json:"method"`
	BankAccountHolder string `json:"bank_account_holder"`
	BankName          string `json:"bank_name"`
	BankBranchCode    string `json:"bank_branch_code"`
	BankAccountNumber string `json:"bank_account_number"`
	BankAccountType   string `json:"bank_account_type"`
}

type BankNames struct {
	BankName string `json:"bank_name,omitempty"`
	Label    string `json:"label,omitempty"`
}

type Refund struct {
	AvailableBalance int                 `json:"available_balance"`
	Transactions     []RefundTransaction `json:"transactions"`
}

type RefundTransaction struct {
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
	Type   string    `json:"type"`
}
