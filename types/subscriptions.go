package types

import (
	"time"

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

func (s *Subscription) Copy(data map[string]interface{}) {
	s.Amount = int(data["amount"].(float64))
	s.Cycles = int(data["cycles"].(float64))
	s.CyclesComplete = int(data["cycles_complete"].(float64))
	s.Frequency = int(data["frequency"].(float64))
	s.RunDate = timeutils.FromStandardString(data["run_date"].(string))
	s.Status = int(data["status"].(float64))
	s.StatusReason = data["status_reason"].(string)
	s.StatusText = data["status_text"].(string)
	s.Token = data["token"].(string)
}
