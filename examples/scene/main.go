// Code generated by dol build. Only Generate by scenes if not existed.
// source: main.go

package main

import (
	// "github.com/go-sql-driver/mysql" init
	_ "github.com/go-sql-driver/mysql"
	// "scene/app" init
	"scene/api"
)

//go:generate dolphin build
func main() {
	api.Run()
}
