// Code generated by dol build. Only Generate by tools if not existed.
// source: main.go

package main

import (
	// github.com/2637309949/dolphin/platform/conf init
	_ "github.com/2637309949/dolphin/platform/conf"
	// "github.com/mattn/go-sqlite3" init
	_ "github.com/mattn/go-sqlite3"
	// github.com/2637309949/dolphin/platform/app init
	"github.com/2637309949/dolphin/platform/app"
)

//go:generate dolphin build sqltpl
func main() {
	app.Run()
}
