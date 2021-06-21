// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"sync"

	"github.com/spf13/viper"
)

// MSeti model template
type MSeti interface {
	Add(interface{}, ...string)
	Get(func(string, interface{}) bool) interface{}
	ForEach(func(string, interface{}) error, ...string) error
	Name(func(string) bool) []string
	Release()
}

// MSet struct
type MSet struct {
	lock *sync.RWMutex
	m    map[string][]interface{}
}

// Get defined get models
func (s *MSet) Get(cb func(string, interface{}) bool) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for name, m := range s.m {
		for index := 0; index < len(m); index++ {
			if cb(name, m[index]) {
				return m[index]
			}
		}
	}
	return nil
}

// Add defined add models
func (s *MSet) Add(m interface{}, n ...string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	ns := viper.GetString("app.name")
	if len(n) > 0 {
		ns = n[0]
	}
	s.m[ns] = append(s.m[ns], m)
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
func (s *MSet) ForEach(cb func(name string, m interface{}) error, names ...string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for name, m := range s.m {
		for index := 0; index < len(m); index++ {
			if len(names) > 0 && names[0] == name {
				err := cb(name, m[index])
				if err != nil {
					return err
				}
			} else if len(names) == 0 {
				err := cb(name, m[index])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Release defined release models
func (s *MSet) Release() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = nil
}
