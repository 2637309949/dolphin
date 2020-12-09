package math

import (
	"fmt"
	"math"
	"strconv"
)

// ParseInt64 defined
func ParseInt64(target interface{}, num ...int64) int64 {
	item, err := strconv.ParseInt(fmt.Sprintf("%v", target), 10, 64)
	if err != nil && len(num) > 0 {
		return num[0]
	}
	return item
}

// ParseUInt64 defined
func ParseUInt64(target interface{}, num ...uint64) uint64 {
	item, err := strconv.ParseUint(fmt.Sprintf("%v", target), 10, 64)
	if err != nil && len(num) > 0 {
		return num[0]
	}
	return item
}

// ParseBool defined
func ParseBool(target interface{}, num ...bool) bool {
	item, err := strconv.ParseBool(fmt.Sprintf("%v", target))
	if err != nil && len(num) > 0 {
		return num[0]
	}
	return item
}

// ParseFloat64 defined
func ParseFloat64(target interface{}, num ...float64) float64 {
	item, err := strconv.ParseFloat(fmt.Sprintf("%v", target), 64)
	if err != nil && len(num) > 0 {
		return num[0]
	}
	return item
}

// Round defined
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// SafeMul defined
func SafeMul(a, b uint) uint {
	c := a * b
	if a > 1 && b > 1 && c/b != a {
		return 0
	}
	return c
}
