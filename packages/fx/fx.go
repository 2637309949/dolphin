// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fx

import (
	"reflect"

	"github.com/thoas/go-funk"
	"go.uber.org/fx"
)

// FX defined extend fx
type FX struct {
	*fx.App
	lifecycle Lifecycle
	providers []interface{}
	invokers  []interface{}
}

func (f *FX) provide(constructors ...interface{}) {
	constructors = funk.Map(constructors, func(constructor interface{}) interface{} {
		fnV := reflect.ValueOf(constructor)
		if fnV.Kind() != reflect.Func {
			fnV = reflect.MakeFunc(
				reflect.FuncOf([]reflect.Type{}, []reflect.Type{reflect.TypeOf(constructor)}, false),
				func(args []reflect.Value) (results []reflect.Value) {
					return []reflect.Value{reflect.ValueOf(constructor)}
				},
			)
			return fnV.Interface()
		}
		return constructor
	}).([]interface{})
	f.providers = append(f.providers, constructors...)
}

func (f *FX) invoke(funcs ...interface{}) {
	f.invokers = append(f.invokers, funcs...)
}

func (f *FX) build() {
	// disable fx logger
	f.App = fx.New(fx.NopLogger, fx.Provide(f.providers...), fx.Invoke(f.invokers...))
}

// NewFX defined new FX
func NewFX() *FX {
	fx := &FX{lifecycle: &lifecycleWrapper{}}
	return fx
}
