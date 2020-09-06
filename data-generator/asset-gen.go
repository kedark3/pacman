package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	// Temporarily created a /static dir and moved all desired contents in there
	// so that during compilation I can use it to create assets_vfsdata.go
	// then ran `go build data-generator/asset-gen.go`
	// That created a assets_vfsdata.go in root of repo which I then used in
	// conjunction with pacman.go to build ./pacman executable that
	// has all static content pre-pkgd
	var fs http.FileSystem = http.Dir("./static")
	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}
}
