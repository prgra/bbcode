package bbcode

import (
	"testing"

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
		`[b][color=red]Внимание![/color][/b]!Заказ номер [i][color=purple]xxxx[/color][/i]Сделал депозит с [color=green][b]'EUR'[/b][/color] и были пополнения ранее с [i]'Umob_Ivory_Coast'[/i] !аккаунт [color=blue]xxxx[/color] [i][color=navy]'Oke Akintunde'[/color][/i]На сумму [color=maroon]1.52 EUR[/color]Страна аккаунта: Кот-д'Ивуар`: `Внимание!!Заказ номер xxxxСделал депозит с 'EUR' и были пополнения ранее с 'Umob_Ivory_Coast' !аккаунт xxxx 'Oke Akintunde'На сумму 1.52 EURСтрана аккаунта: Кот-д'Ивуар`,
	}
	for k, v := range table {
		s := Parse(k)
		if s.NewString != v {
			t.Errorf("TestParse error for : %s\ngot:  %s\nneed: %s\nDebug info: \n%s", k, s.NewString, v, spew.Sdump(s))

		}
	}
}

func TestPositions(t *testing.T) {
	s := Parse("[b]b[/b][u]u[/u]")
	spew.Dump(s)
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
