package types

type Status int

const (
	Active Status = iota + 1
	Cancelled
	Paused
	Complete
	InReview
	Failed
	SystemPaused
)

func (s Status) String() string {
	return [...]string{"Active", "Cancelled", "Paused", "Complete", "InReview", "Failed", "SystemPaused"}[s]
}
