package bbcode

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// BBCodes List of codes result for Parse function
type BBCodes struct {
	Original  string
	NewString string
	BBCodes   []BBCode
}

// BBCode struct
type BBCode struct {
	OriginalStart int
	OriginalEnd   int
	Pos           int
	Len           int
	Param         string
	Name          string
	IsClose       bool
	IsValid       bool
}

var validCodes map[string]bool = map[string]bool{
	"b":     true,
	"i":     true,
	"u":     true,
	"user":  true,
	"code":  true,
	"quote": true,
}

// Parse string for valid BBCode and return list of parsed values and newstring
func Parse(s string) (b BBCodes) {
	var tag BBCode
	start := false
	b.Original = s
	for i, r := range []rune(s) {
		if string(r) == "[" {
			start = true
			tag.Name = ""
			tag.OriginalStart = i + 1
			tag.Pos = tag.OriginalStart
			continue
		}

		if string(r) == "]" {
			start = false
			tag.OriginalEnd = i + 1
			tag.Pos = tag.OriginalStart
			b.BBCodes = append(b.BBCodes, tag)
			tag = BBCode{}
			continue
		}

		if start {
			tag.Name += string(r)
		}
	}

	// Ищем закрыте и выставляем флаг
	for i := range b.BBCodes {
		if strings.HasPrefix(b.BBCodes[i].Name, "/") {
			b.BBCodes[i].IsClose = true
			b.BBCodes[i].Name = strings.Replace(b.BBCodes[i].Name, "/", "", 1)
		}
		if strings.Index(b.BBCodes[i].Name, "=") != -1 {
			kv := strings.Split(b.BBCodes[i].Name, "=")
			if len(kv) == 2 {
				b.BBCodes[i].Name = kv[0]
				b.BBCodes[i].Param = kv[1]
			}
		}
	}

	// Проходимся, ищем пары
	for i := range b.BBCodes {
		if !b.BBCodes[i].IsClose {
			for j := i; j < len(b.BBCodes); j++ {
				if strings.ToLower(b.BBCodes[i].Name) == strings.ToLower(b.BBCodes[j].Name) &&
					b.BBCodes[j].IsClose {
					b.BBCodes[i].IsValid = validCodes[b.BBCodes[i].Name]
					b.BBCodes[j].IsValid = validCodes[b.BBCodes[i].Name]
					b.BBCodes[i].Len = b.BBCodes[j].OriginalStart - b.BBCodes[i].OriginalEnd - 1
					continue
				}
			}
		}
	}

	// tag code finder and turn off valid flag
	inCode := false
	for i := range b.BBCodes {
		if strings.ToLower(b.BBCodes[i].Name) == "code" && b.BBCodes[i].IsValid && !inCode {
			inCode = true
			continue
		}
		if strings.ToLower(b.BBCodes[i].Name) == "code" && b.BBCodes[i].IsClose {
			inCode = false
		}
		if inCode {
			b.BBCodes[i].IsValid = false
		}
	}

	// Make new string, cut untervals reverse
	spew.Dump(b)
	b.NewString = b.Original
	for i := 0; i < len(b.BBCodes); i++ {
		if b.BBCodes[i].IsValid {
			l := b.BBCodes[i].OriginalEnd - b.BBCodes[i].OriginalStart
			fmt.Printf("cut N%d tag %s %s %d-%d\n", i+1, b.BBCodes[i].Name, b.NewString, b.BBCodes[i].Pos, b.BBCodes[i].Pos+l)
			for j := i + 1; j < len(b.BBCodes); j++ {
				// fmt.Printf("tags N%d shift, %s %d on %d, closed  %v\n", j+1, b.BBCodes[j].Name, b.BBCodes[j].Pos, b.BBCodes[i].OriginalEnd-b.BBCodes[i].OriginalStart, b.BBCodes[i].IsClose)
				fmt.Printf("shi N%d POS %d->%d\n", j+1, b.BBCodes[j].Pos, b.BBCodes[j].Pos-(l+1))
				b.BBCodes[j].Pos -= (l + 1)
			}
			b.NewString = cutString(b.NewString, b.BBCodes[i].Pos, b.BBCodes[i].Pos+l)
			fmt.Println(b.NewString)
		}
	}
	return
}

func cutString(s string, f int, t int) (res string) {
	for i, r := range []rune(s) {
		if i+1 < f || i+1 > t {
			res += string(r)
		}
	}
	return
}
