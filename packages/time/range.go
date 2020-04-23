package time

import "time"

// Iter is a timerange iterator.
type Iter struct {
	start    time.Time
	end      time.Time
	interval time.Duration
	current  time.Time
	index    int
}

// NewRange an Iter.
func NewRange(start time.Time, end time.Time, interval time.Duration) *Iter {
	return &Iter{
		start:    start,
		end:      end,
		interval: interval,
		current:  start,
		index:    0,
	}
}

// Next scans for the next time.
func (i *Iter) Next() bool {
	var next time.Time
	if i.index == 0 {
		next = i.current
	} else {
		next = i.current.Add(i.interval)
	}

	if i.end.Equal(next) || i.end.After(next) {
		i.current = next
		i.index++
		return true
	}
	return false
}

// Current returns the latest unscanned time.
func (i *Iter) Current() time.Time {
	return i.current
}
