package math

import (
	"strconv"
	"time"
)

// MustInt64 defined TODO
func MustInt64(x string) int64 {
	i, _ := strconv.ParseInt(x, 10, 64)
	return i
}

// MustInt32 defined TODO
func MustInt32(x string) int32 {
	i, _ := strconv.ParseInt(x, 10, 16)
	return int32(i)
}

// MustInt16 defined TODO
func MustInt16(x string) int16 {
	i, _ := strconv.ParseInt(x, 10, 16)
	return int16(i)
}

// MustInt8 defined TODO
func MustInt8(x string) int8 {
	i, _ := strconv.ParseInt(x, 10, 8)
	return int8(i)
}

// MustInt defined TODO
func MustInt(x string) int {
	i, _ := strconv.ParseInt(x, 10, 0)
	return int(i)
}

// MustFloat32 defined TODO
func MustFloat32(x string) float32 {
	i, _ := strconv.ParseFloat(x, 32)
	return float32(i)
}

// MustFloat64 defined TODO
func MustFloat64(x string) float64 {
	i, _ := strconv.ParseFloat(x, 64)
	return float64(i)
}

// MustBool defined TODO
func MustBool(x string) bool {
	i, _ := strconv.ParseBool(x)
	return i
}

// MustTime defined TODO
func MustTime(x string, layout string) time.Time {
	i, _ := time.Parse(layout, x)
	return i
}
