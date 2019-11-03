package timeutils

import "time"

const FORMAT = "2006-01-02T15:04:05-07:00"

func ToStandardString(time time.Time) string {
	return time.Format(FORMAT)
}

func FromStandardString(text string) time.Time {
	t, _ := time.Parse(FORMAT, text)
	return t
}