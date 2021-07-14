package maps

// MustMap Defined TODO
type MustMap map[string]interface{}

// MustInt64 defined TODO
func (q *MustMap) MustInt64(key string) int64 {
	i, _ := (*q)[key].(int64)
	return i
}

// MustInt defined TODO
func (q *MustMap) MustInt(key string) int {
	i, _ := (*q)[key].(int)
	return i
}

// MustFloat32 defined TODO
func (q *MustMap) MustFloat32(key string) float32 {
	i, _ := (*q)[key].(float32)
	return i
}

// MustFloat32 defined TODO
func (q *MustMap) MustFloat64(key string) float64 {
	i, _ := (*q)[key].(float64)
	return i
}

// MustBool defined TODO
func (q *MustMap) MustBool(key string) bool {
	i, _ := (*q)[key].(bool)
	return i
}

// MustString defined TODO
func (q *MustMap) MustString(key string) string {
	i, _ := (*q)[key].(string)
	return i
}
