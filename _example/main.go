package main

import (
	"github.com/prgra/bbcode"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := bbcode.Parse("[b]hello[/b][code]test[/code]")
	spew.Dump(a)
}
