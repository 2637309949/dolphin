// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fx

// Option defined implement of option
type (
	Option func(*Dol) interface{}
)

// option defined implement of option
func (o Option) apply(r *Dol) { o(r) }

// InvokeOptions defined invokes
func InvokeOptions(invokers ...interface{}) Option {
	return Option(func(d *Dol) interface{} {
		for _, v := range invokers {
			d.fx.invoke(v)
		}
		return d
	})
}

// ProviderOptions defined providers
func ProviderOptions(providers ...interface{}) Option {
	return Option(func(d *Dol) interface{} {
		for _, v := range providers {
			d.fx.provide(v)
		}
		return d
	})
}
