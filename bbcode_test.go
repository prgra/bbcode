package bbcode

import (
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestParse(t *testing.T) {
	table := map[string]string{
		"[txt=ololo][code][i]II[/i][/zmg]Bo[/b]llo a😪a 🍻 ss": "[txt=ololo][code]II[/zmg]Bo[/b]llo a😪a 🍻 ss",
		"[code][code][b]sdfs--[/code]":                       "[code][b]sdfs--",
		"[b][b][b]b[/b]":                                     "[b][b]b",
		"[b][b][i][b]b[/b]":                                  "[b][b][i]b",
		"[i]ii[/i][b]b[/b][/b][/b]":                          "iib[/b][/b]",
		"[i]1[/i][i]z[/i]":                                   "1z",
		"[color=red]1[color=green]gr[/color]red[/color]":     "1grred",
		"[Color=red]1[color=green]gr[/Color]red[/color]":     "1grred",
		"[[b]][[/b]]":                                        "[][]",
		"[b]b[/b][u]u[/u]":                                   "bu",
		"[B]b[/b][U]u[/U]":                                   "bu",
		"[color=black]black[code][/code][/a][/color][ssss]":  "black[/a][ssss]",
	}
	for k, v := range table {
		s := Parse(k)
		if s.NewString != v {
			t.Errorf("TestParse error for : %s\ngot:  %s\nneed: %s\nDebug info: \n%s", k, s.NewString, v, spew.Sdump(s))

		}
	}
}

func BenchmarkParse(b *testing.B) {
	table := []string{
		"[txt=ololo][code][i]II[/i][/zmg]Bo[/b]llo a😪a 🍻 ss",
		"[code][code][b]sdfs--[/code]",
		"[b][b][b]b[/b]",
		"[b][b][i][b]b[/b]",
		"[i]ii[/i][b]b[/b][/b][/b]",
		"[i]1[/i][i]z[/i]",
		"[color=red]1[color=green]gr[/color]red[/color]",
	}
	rand.Seed(time.Now().UnixNano())
	j := 0
	for i := 0; i < b.N; i++ {
		j++
		if j > len(table)-1 {
			j = 0
		}
		Parse(table[j])
	}
}

func TestPositions(t *testing.T) {
	s := Parse("[b]b[/b][u]u[/u]")
	if len(s.BBCodes) != 4 {
		t.Error("count bbcode err need 4")
	}
	if s.BBCodes[0].Pos != 1 || s.BBCodes[2].Pos != 2 {
		t.Error("poses err")
	}
	s = Parse("[b]👩test[/b][i]👩‍👩‍👦‍👦ing[/i]123")
	if s.NewString != "👩test👩‍👩‍👦‍👦ing123" {
		t.Errorf("TestPositions new string: got:  %s\nneed: 👩test \nDebug info: \n%v", s.NewString, spew.Sdump(s))
	}
	if s.BBCodes[0].Len != 6 {
		t.Errorf("TestPositions error: got:  %d\nneed: 6\nDebug info: \n%v", s.BBCodes[0].Len, spew.Sdump(s))
	}
	if s.BBCodes[1].Pos != 7 {
		t.Errorf("TestPositions error: got:  %d\nneed: 7\nDebug info: \n%v", s.BBCodes[1].Pos, spew.Sdump(s))
	}
}

func TestUrl(t *testing.T) {
	s := Parse("[url=http://google.com?q=123]test ok[/url]")
	if s.NewString != "test ok" {
		t.Errorf("wrong text need 'test ok' got %s\nDebug info: \n%v", s.NewString, spew.Sdump(s))
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
