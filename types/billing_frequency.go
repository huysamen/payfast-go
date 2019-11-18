package types

type BillingFrequency int

const (
	Monthly BillingFrequency = iota + 3
	Quarterly
	Biannual
	Annual
)

func (bf BillingFrequency) String() string {
	return [...]string{"Monthly", "Quarterly", "Biannual", "Annual"}[bf]
}
