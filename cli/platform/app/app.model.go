package app

// MSeti model template
type MSeti interface {
	Add(m interface{})
	Get(func(interface{}) bool) interface{}
	ForEach(func(interface{}))
}

// MSets struct
type MSets struct {
	m []interface{}
}

// Get defined get models
func (s *MSets) Get(cb func(interface{}) bool) interface{} {
	var model interface{}
	s.ForEach(func(m interface{}) {
		if cb(m) {
			model = m
		}
	})
	return model
}

// Add defined add models
func (s *MSets) Add(m interface{}) {
	s.m = append(s.m, m)
}

// ForEach defined foreach models
func (s *MSets) ForEach(cb func(m interface{})) {
	for index := 0; index < len(s.m); index++ {
		cb(s.m[index])
	}
}
