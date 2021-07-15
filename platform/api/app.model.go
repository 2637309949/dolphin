// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"sync"

	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/spf13/viper"
)

// Table defined TODO
type Table interface {
	TableName() string
}

// ModelSetter model template
type ModelSetter interface {
	Add(Table, ...string)
	ByName(string) []Table
	ByNotName(string) []Table
	NameSpaces(...string) (names []string)
	Release()
}

// MSet struct
type defaultModelSetter struct {
	lock *sync.RWMutex
	m    map[string][]Table
}

// Add defined add models
func (s *defaultModelSetter) NameSpaces(n ...string) (names []string) {
	for k := range s.m {
		if slice.StrSliceContains(n, k) {
			continue
		}
		names = append(names, k)
	}
	return
}

// Add defined add models
func (s *defaultModelSetter) Add(m Table, n ...string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(n) == 0 {
		n = []string{viper.GetString("app.name")}
	}
	s.m[n[0]] = append(s.m[n[0]], m)
}

// Name defined add models
func (s *defaultModelSetter) ByName(name string) []Table {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.m[name]
}

// Name defined add models
func (s *defaultModelSetter) ByNotName(name string) (m []Table) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for k, v := range s.m {
		if k != name {
			m = append(m, v...)
		}
	}
	return
}

// Release defined release models
func (s *defaultModelSetter) Release() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = nil
}
