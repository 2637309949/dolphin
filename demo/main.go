// Code generated by dol build. Only Generate by tools if not existed.
// source: main.go

package main

import (
	_ "demo/app"

	"github.com/2637309949/dolphin/packages/fx/cli"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate dolphin build
func main() {
	cli.Run()
}
