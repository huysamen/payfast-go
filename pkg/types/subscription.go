package types

import (
	"time"
)

type Subscription struct {
	Amount         int       `json:"amount"`
	Cycles         int       `json:"cycles"`
	CyclesComplete int       `json:"cycles_complete"`
	Frequency      int       `json:"frequency"`
	RunDate        time.Time `json:"run_date"`
	Status         int       `json:"status"`
	StatusReason   string    `json:"status_reason"`
	StatusText     string    `json:"status_text"`
	Token          string    `json:"token"`
}
