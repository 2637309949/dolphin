package null

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	zeroTime   = "1970-01-01 00:00:00"
)

func Parse(x string) (time.Time, error) {
	if strings.Index(x, "T") > 0 {
		x = strings.Replace(x, "T", " ", 1)
	}
	switch l := len(x); true {
	case l > 19:
		x = x[0:19]
		if strings.Index(x, "000") == 0 {
			x = zeroTime
		}
	case l == 10:
		x += " 00:00:00"
	case l == 13:
		x += ":00:00"
	case l == 16:
		x += ":00"
	}
	return time.ParseInLocation(timeFormat, x, time.Local)
}

// ToDB implements json.Marshaler.
// It will encode null if this time is null.
func (t Time) ToDB() ([]byte, error) {
	return t.MarshalJSON()
}

// FromDB implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (t Time) FromDB(data []byte) error {
	data, err := json.Marshal(string(data))
	if err != nil {
		return err
	}
	return t.UnmarshalJSON(data)
}
