// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/2637309949/dolphin/srv"
)

var _dol = srv.New()

// Use defined use plugin
func Use(opts ...srv.Option) {
	_dol.Use(opts...)
}

// Invoke defined use plugin
func Invoke(invokers ...interface{}) error {
	_dol.Use(srv.InvokeOptions(invokers...))
	return nil
}

// Provider defined use plugin
func Provider(providers ...interface{}) error {
	_dol.Use(srv.ProviderOptions(providers...))
	return nil
}

// Run defined run application
func Run() error {
	return _dol.Run()
}
