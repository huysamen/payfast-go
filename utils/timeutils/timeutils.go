package timeutils

import "time"

const FullFormat = "2006-01-02T15:04:05-07:00"
const CsvFormat = "2006-01-02 15:04:05"
const DayFormat = "2006-01-02"
const MonthFormat = "2006-01"

func ToStandardString(time time.Time) string {
	return time.Format(FullFormat)
}

func FromStandardString(text string) time.Time {
	t, _ := time.Parse(FullFormat, text)
	return t
}

func FromCsvString(text string) time.Time {
	t, _ := time.Parse(CsvFormat, text)
	return t
}

func ToDayString(time time.Time) string {
	return time.Format(DayFormat)
}

func ToMonthString(time time.Time) string {
	return time.Format(MonthFormat)
}
