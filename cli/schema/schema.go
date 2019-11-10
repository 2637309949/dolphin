// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package schema

// Success struct
type Success struct {
	Common
	Ref string
}

// Failure struct
type Failure struct {
	Common
	Ref string
}

// Return struct
type Return struct {
	Common
	Success *Success
	Failure *Failure
}

// Param struct
type Param struct {
	Common
	Name string `validate:"required"`
	Desc string `validate:"required"`
	Ref  string
	Type string
}

// API struct
type API struct {
	Common
	Name     string `validate:"required"`
	Desc     string `validate:"required"`
	Version  string
	Function string
	Table    string
	Method   string
	Params   []*Param
	Return   *Return
}

// Controller struct
type Controller struct {
	Common
	Name      string `validate:"required"`
	Desc      string `validate:"required"`
	SkipLogin bool
	APIS      []*API
}

// Prop struct
type Prop struct {
	Common
	Name string `validate:"required"`
	Desc string `validate:"required"`
	Type string
}

// Bean struct
type Bean struct {
	Common
	Name     string `validate:"required"`
	Desc     string `validate:"required"`
	Packages string
	Props    []*Prop
}

// Column struct
type Column struct {
	Common
	Name string `validate:"required"`
	Desc string `validate:"required"`
	Type string
	Xorm string
}

// Table struct
type Table struct {
	Common
	Name     string `validate:"required"`
	Desc     string `validate:"required"`
	Packages string
	Columns  []*Column
}

// Application struct
type Application struct {
	Common
	Name        string `validate:"required"`
	Desc        string `validate:"required"`
	PackageName string `validate:"required"`
	Controllers []*Controller
	Beans       []*Bean
	Tables      []*Table
}
