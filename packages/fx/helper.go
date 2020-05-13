// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fx

// Invoke defined use plugin
func Invoke(invokers ...interface{}) error {
	d.Use(InvokeOptions(invokers...))
	return nil
}

// Provider defined use plugin
func Provider(providers ...interface{}) error {
	d.Use(ProviderOptions(providers...))
	return nil
}
