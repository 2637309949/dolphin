package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../assets")

func main() {
	if err := os.MkdirAll("../dist", os.ModePerm); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	err := vfsgen.Generate(Assets, vfsgen.Options{
		Filename:     "../dist/assets.go",
		PackageName:  "dist",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
