package app

// MSeti model template
type MSeti interface {
	Add(string, interface{})
	Get(func(string, interface{}) bool) interface{}
	ForEach(func(string, interface{}))
	Name(func(string) bool) []string
}

// MSet struct
type MSet struct {
	m map[string][]interface{}
}

// Get defined get models
func (s *MSet) Get(cb func(string, interface{}) bool) interface{} {
	var model interface{}
	s.ForEach(func(name string, m interface{}) {
		if cb(name, m) {
			model = m
		}
	})
	return model
}

// Add defined add models
func (s *MSet) Add(n string, m interface{}) {
	s.m[n] = append(s.m[n], m)
}

// Name defined add models
func (s *MSet) Name(cb func(string) bool) (m []string) {
	for k := range s.m {
		if cb(k) {
			m = append(m, k)
		}
	}
	return
}

// ForEach defined foreach models
func (s *MSet) ForEach(cb func(name string, m interface{})) {
	for name, m := range s.m {
		for index := 0; index < len(m); index++ {
			cb(name, m[index])
		}
	}
}
