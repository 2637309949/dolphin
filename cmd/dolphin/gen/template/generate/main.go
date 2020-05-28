package main

import (
	"log"
	"net/http"

	"github.com/2637309949/dolphin/packages/vfsgen"
)

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../assets")

func main() {
	err := vfsgen.Generate(Assets, vfsgen.Options{
		Filename:     "../assets.go",
		PackageName:  "template",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
