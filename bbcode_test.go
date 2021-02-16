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
		"[color=black]black[code][a]ololo2[/code][/color]":   "black[a]ololo2",
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
	if s.BBCodes[0].Len != 5 {
		t.Errorf("TestPositions error: got:  %d\nneed: 5\nDebug info: \n%v", s.BBCodes[0].Len, spew.Sdump(s))
	}
	if s.BBCodes[1].Pos != 6 {
		t.Errorf("TestPositions error: got:  %d\nneed: 4\nDebug info: \n%v", s.BBCodes[1].Pos, spew.Sdump(s))
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

func TestMakeURLs(t *testing.T) {
	a := "[url=http://google.com]http://bing.com[/url]https://mail.com"
	bb := Parse(a)
	bb.MakeURLs()
	if len(bb.BBCodes) != 4 {
		t.Errorf("wrong count of bbcodes got:%d, need: 4", len(bb.BBCodes))
	}
	if bb.BBCodes[3].Name != "url" {
		t.Errorf("wrong name bbcode got:%s, need: url", bb.BBCodes[3].Name)
	}
	bb = Parse("[url=http://ya.ru/q=Tv_vT]asda[/url]")
	if bb.NewString != "asda" {
		t.Errorf("wrong new string: need 'asda', got '%s'", bb.NewString)
	}
	if bb.BBCodes[0].Param != "http://ya.ru/q=Tv_vT" {
		t.Errorf("wrong Param: need 'asda', got '%s'", bb.NewString)
	}

}

func TestUtfLens(t *testing.T) {
	str := "Напоминание[color=gray]🔆️[/color] [user=490]Name[/user]"
	bb := Parse(str)
	if bb.BBCodes[0].Len != 1 {
		t.Errorf("wrong Len 0: need 1, got '%d'", bb.BBCodes[0].Len)
	}
	if bb.BBCodes[0].Pos != 12 {
		t.Errorf("wrong Pos 0: need 12, got '%d'", bb.BBCodes[0].Len)
	}
	if bb.BBCodes[2].Pos != 14 {
		t.Errorf("wrong Pos 2: need 14, got '%d'", bb.BBCodes[2].Pos)
	}
	if bb.BBCodes[3].Pos != 18 {
		t.Errorf("wrong Pos 3: need 18, got '%d'", bb.BBCodes[3].Pos)
	}
}
