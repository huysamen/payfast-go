package types

import "time"

type Numeric struct {
	Valid bool
	Value int
}

type AlphaNumeric struct {
	Valid bool
	Value string
}

type Time struct {
	Valid bool
	Value time.Time
}

type Bool struct {
	Valid bool
	Value bool
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
