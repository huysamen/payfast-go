package types

import (
	"time"

	"github.com/huysamen/payfast-go/utils/copyutils"
	"github.com/huysamen/payfast-go/utils/timeutils"
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

func (s *Subscription) Copy(data map[string]any) {
	s.Amount = copyutils.CopyInt(data, "amount")
	s.Cycles = copyutils.CopyInt(data, "cycles")
	s.CyclesComplete = copyutils.CopyInt(data, "cycles_complete")
	s.Frequency = copyutils.CopyInt(data, "frequency")
	s.RunDate = timeutils.FromStandardString(data["run_date"].(string))
	s.Status = copyutils.CopyInt(data, "status")
	s.StatusReason = copyutils.CopyString(data, "status_reason")
	s.StatusText = copyutils.CopyString(data, "status_text")
	s.Token = copyutils.CopyString(data, "token")
}
