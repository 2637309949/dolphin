package app

import "github.com/2637309949/dolphin/packages/viper"

// MSeti model template
type MSeti interface {
	Add(interface{}, ...string)
	Get(func(string, interface{}) bool) interface{}
	ForEach(func(string, interface{}), ...string)
	Name(func(string) bool) []string
	Release()
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
func (s *MSet) Add(m interface{}, n ...string) {
	if len(n) > 0 {
		s.m[n[0]] = append(s.m[n[0]], m)
	} else {
		s.m[viper.GetString("app.name")] = append(s.m[viper.GetString("app.name")], m)
	}
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
func (s *MSet) ForEach(cb func(name string, m interface{}), names ...string) {
	for name, m := range s.m {
		for index := 0; index < len(m); index++ {
			if len(names) > 0 && names[0] == name {
				cb(name, m[index])
			} else if len(names) == 0 {
				cb(name, m[index])
			}
		}
	}
}

// Release defined release models
func (s *MSet) Release() {
	s.m = nil
}
