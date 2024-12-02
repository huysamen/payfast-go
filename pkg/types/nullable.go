package types

import (
	"strconv"
	"time"

	"github.com/huysamen/payfast-go/internal/utils/timeutils"
)

type Numeric struct {
	Valid bool
	Value int
}

func (n Numeric) Val() any {
	if !n.Valid {
		return 0
	}

	return n.Value
}

func (n Numeric) String() string {
	if !n.Valid {
		return ""
	}

	return strconv.Itoa(n.Value)
}

type AlphaNumeric struct {
	Valid bool
	Value string
}

func (a AlphaNumeric) Val() any {
	if !a.Valid {
		return ""
	}

	return a.Value
}

func (a AlphaNumeric) String() string {
	if !a.Valid {
		return ""
	}

	return a.Value
}

type Time struct {
	Valid bool
	Value time.Time
}

func (t Time) Val() any {
	if !t.Valid {
		return time.Time{}
	}

	return t.Value
}

func (t Time) String() string {
	if !t.Valid {
		return ""
	}

	return timeutils.ToDayString(t.Value)
}

type Bool struct {
	Valid bool
	Value bool
}

func (b Bool) Val() any {
	if !b.Valid {
		return false
	}

	return b.Value
}

func (b Bool) String() string {
	if !b.Valid {
		return ""
	}

	if b.Value {
		return "1"
	}

	return "0"
}

func NewNumeric(value int) Numeric {
	return Numeric{
		Valid: true,
		Value: value,
	}
}

func NewAlphaNumeric(value string) AlphaNumeric {
	return AlphaNumeric{
		Valid: true,
		Value: value,
	}
}

func NewTime(value time.Time) Time {
	return Time{
		Valid: true,
		Value: value,
	}
}

func NewBool(value bool) Bool {
	return Bool{
		Value: true,
		Valid: value,
	}
}
