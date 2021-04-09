package time

import (
	"time"
)

// A Time represents an instant in time with nanosecond precision.
type Time struct {
	time.Time
}

// Now returns the current local time.
func Now() Time {
	return Time{time.Now()}
}

// Value returns t as a Unix time
func (t Time) Value() time.Time {
	return t.Time
}

// UnixMilli returns t as a Unix time
func (t Time) UnixMilli() int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// StartOfWeek defined
func (t Time) StartOfWeek() time.Time {
	return t.AddDate(0, 0, int(time.Monday-t.Weekday()))
}

// EndOfWeek defined
func (t Time) EndOfWeek() time.Time {
	return t.AddDate(0, 0, int(time.Sunday-t.Weekday())+7)
}

// StartOfMonth defined
func (t Time) StartOfMonth() time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

// EndOfMonth defined
func (t Time) EndOfMonth() time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

// Parse defined
func Parse(layout string, value string) (Time, error) {
	te, err := time.Parse(layout, value)
	return Time{te}, err
}
