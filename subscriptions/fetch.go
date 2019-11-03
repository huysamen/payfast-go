package subscriptions

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/huysamen/payfast-go/types"
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

func (c *Client) Fetch(subscriptionID string) (*Subscription, error) {
	body, err := c.get(strings.ReplaceAll(fetchPath, "__sid__", subscriptionID))
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response)

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return nil, err
	}

	if rsp.Code == 200 {
		sub := &Subscription{}
		sub.Copy(rsp.Data.Response.(map[string]interface{}))

		return sub, nil
	}

	return nil, nil
}
