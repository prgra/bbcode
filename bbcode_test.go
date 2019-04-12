package bbcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
		"[[b]][[/b]]":                                        "[][]",
	}
	for k, v := range table {
		s := Parse(k)
		if s.NewString != v {
			t.Errorf("TestParse error for : %s\ngot:  %s\nneed: %s\nDebug info: \n%v", k, s.NewString, v, s)

		}
	}
}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
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
	if s.NewString != "bu" {
		t.Fail()
	}
	if len(s.BBCodes) != 4 {
		t.Fail()
	}
	if s.BBCodes[0].Pos != 1 || s.BBCodes[2].Pos != 2 {
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
