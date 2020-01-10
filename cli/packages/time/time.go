package time

import (
	"time"
	gt "time"
)

// A Time represents an instant in time with nanosecond precision.
type Time struct {
	gt.Time
}

// Now returns the current local time.
func Now() Time {
	return Time{gt.Now()}
}

// UnixMilli returns t as a Unix time
func (t Time) UnixMilli() int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
