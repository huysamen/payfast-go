package types

const (
	Active int = iota + 1
	Cancelled
	Paused
	Complete
	InReview
	Failed
	SystemPaused
)
