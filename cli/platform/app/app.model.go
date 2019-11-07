package app

// MSeti model template
type MSeti interface {
	Add(string, interface{})
	Get(func(string, interface{}) bool) interface{}
	ForEach(func(string, interface{}))
}

// MSets struct
type MSets struct {
	m map[string][]interface{}
}

// Get defined get models
func (s *MSets) Get(cb func(string, interface{}) bool) interface{} {
	var model interface{}
	s.ForEach(func(name string, m interface{}) {
		if cb(name, m) {
			model = m
		}
	})
	return model
}

// Add defined add models
func (s *MSets) Add(n string, m interface{}) {
	s.m[n] = append(s.m[n], m)
}

// ForEach defined foreach models
func (s *MSets) ForEach(cb func(name string, m interface{})) {
	for name, m := range s.m {
		for index := 0; index < len(m); index++ {
			cb(name, m[index])
		}
	}
}
