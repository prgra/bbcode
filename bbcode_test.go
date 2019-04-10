package bbcode

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParse(t *testing.T) {
	s := Parse("[txt=ololo][code][i]II[/i][/zmg]Bo[/b]llo a😪a 🍻 ss")
	if s.NewString != "[txt=ololo][code]II[/zmg]Bo[/b]llo a😪a 🍻 ss" {
		t.Fail()
	}
	s = Parse("[code][code][b]sdfs--[/code]")
	if s.NewString != "[code][b]sdfs--" {
		t.Fail()
	}
}

func TestPositions(t *testing.T) {
	s := Parse("[b]b[/b][u]u[/u]")
	spew.Dump(s)
	if s.NewString != "[txt=ololo][code]II[/zmg]Bo[/b]llo a😪a 🍻 ss" {
		t.Fail()
	}
}

func TestCutString(t *testing.T) {
	s := "1 2 3 4"
	if cutString(s, 3, 6) != "1 4" {
		t.Fail()
	}
	if cutString("", 3, 6) != "" {
		t.Fail()
	}
	if cutString("asdasdasdasda", 0, 0) != "asdasdasdasda" {
		t.Fail()
	}
}
