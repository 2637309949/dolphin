package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../assets")

func main() {
	err := vfsgen.Generate(Assets, vfsgen.Options{
		Filename:     "../assets/assets.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
