package bbcode

import (
	"strings"
	"unicode/utf8"
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
	CloseFor      int
	OpenFor       int
	Text          string
}

var validCodes = map[string]bool{
	"b":     true,
	"i":     true,
	"u":     true,
	"s":     true,
	"user":  true,
	"code":  true,
	"quote": true,
	"color": true,
	"size":  true,
	"url":   true,
}

var multiCodes = map[string]bool{
	"color": true,
}

// Parse string for valid BBCode and return list of parsed values and newstring
func Parse(s string) (b BBCodes) {
	var i, j int
	var tag BBCode
	start := false
	b.Original = s
	rs := []rune(s)
	for i = range rs {
		tag.CloseFor = -1
		tag.OpenFor = -1
		if string(rs[i]) == "[" {
			for j = i + 1; j < len(rs); j++ {
				if string(rs[j]) == "[" {
					start = false
					break
				}
				if string(rs[j]) == "]" {
					start = true
					tag.Name = ""
					tag.OriginalStart = i + 1
					tag.Pos = tag.OriginalStart
					break
				}
			}
			continue
		}
		if string(rs[i]) == "]" && start {
			start = false
			tag.OriginalEnd = i + 1
			tag.Pos = tag.OriginalStart
			b.BBCodes = append(b.BBCodes, tag)
			tag = BBCode{}
			continue
		}

		if start {
			tag.Name += string(rs[i])
		}
	}
	// Ищем закрыте и выставляем флаг
	for i = range b.BBCodes {
		if strings.HasPrefix(b.BBCodes[i].Name, "/") {
			b.BBCodes[i].IsClose = true
			b.BBCodes[i].Name = strings.ToLower(strings.Replace(b.BBCodes[i].Name, "/", "", 1))
		}
		if strings.Index(b.BBCodes[i].Name, "=") != -1 {
			kv := strings.Split(b.BBCodes[i].Name, "=")
			if len(kv) >= 2 {
				b.BBCodes[i].Name = strings.ToLower(kv[0])
				b.BBCodes[i].Param = strings.Join(kv[1:], "=")
			}
		} else {
			b.BBCodes[i].Name = strings.ToLower(b.BBCodes[i].Name)
		}
		if !b.BBCodes[i].IsClose && len(b.BBCodes) > i+1 {
			b.BBCodes[i].Text, b.BBCodes[i].Len = getRunesRange(b.Original, b.BBCodes[i].OriginalEnd, b.BBCodes[i+1].OriginalStart-1)
		}
	}

	// Проходимся, ищем пары, `c` учитывеат правильный закрывающий тег
	for i = range b.BBCodes {
		if !b.BBCodes[i].IsClose {
			cm := make(map[string]int)
			cm[b.BBCodes[i].Name]++
			for j = i + 1; j < len(b.BBCodes); j++ {
				if b.BBCodes[i].Name == b.BBCodes[j].Name && !b.BBCodes[j].IsClose &&
					!multiCodes[b.BBCodes[i].Name] {
					b.BBCodes[i].IsValid = false
					break
				}
				if !b.BBCodes[j].IsClose {
					cm[b.BBCodes[j].Name]++
				}
				if b.BBCodes[j].IsClose {
					cm[b.BBCodes[j].Name]--
				}
				if b.BBCodes[i].Name == b.BBCodes[j].Name &&
					b.BBCodes[j].IsClose &&
					(cm[b.BBCodes[i].Name] == 0) {
					b.BBCodes[i].IsValid = validCodes[b.BBCodes[i].Name]
					b.BBCodes[j].IsValid = validCodes[b.BBCodes[i].Name]
					b.BBCodes[i].OpenFor = j
					b.BBCodes[j].CloseFor = i
					// b.BBCodes[i].Len = b.BBCodes[j].OriginalStart - b.BBCodes[i].OriginalEnd - 1
					break
				}
			}
		}
	}

	// tag code finder and turn off valid flag
	inCode := false
	for i = range b.BBCodes {
		if b.BBCodes[i].Name == "code" && b.BBCodes[i].IsValid && !inCode {
			inCode = true
			continue
		}
		if b.BBCodes[i].Name == "code" && b.BBCodes[i].IsClose {
			inCode = false
		}
		if inCode {
			b.BBCodes[i].IsValid = false
		}
	}

	// Make new string, cut untervals and shift positions
	b.NewString = b.Original
	for i := range b.BBCodes {
		if b.BBCodes[i].IsValid {
			l := b.BBCodes[i].OriginalEnd - b.BBCodes[i].OriginalStart
			for j = i + 1; j < len(b.BBCodes); j++ {
				b.BBCodes[j].Pos -= l + 1
			}
			b.NewString = cutString(b.NewString, b.BBCodes[i].Pos, b.BBCodes[i].Pos+l)
		}
	}
	ra := []rune(b.NewString)
	for i := range b.BBCodes {
		if b.BBCodes[i].IsValid && !b.BBCodes[i].IsClose {
			p := b.BBCodes[i].OpenFor
			if p < len(b.BBCodes) && p >= 0 {
				b.BBCodes[i].Len = UTF16Count(string(ra[b.BBCodes[i].Pos-1 : b.BBCodes[p].Pos-1]))
				b.BBCodes[i].Text = string(ra[b.BBCodes[i].Pos-1 : b.BBCodes[p].Pos-1])
			}
		}
		// fmt.Printf("%d->%d '%s'   \n", b.BBCodes[i].Pos, UTF16Count(string(ra[0:b.BBCodes[i].Pos-1])), string(ra[0:b.BBCodes[i].Pos]))
		b.BBCodes[i].Pos = UTF16Count(string(ra[0:b.BBCodes[i].Pos-1])) + 1
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

// MakeURLs make url tags from text
func (b *BBCodes) MakeURLs() {
	nstr := b.NewString
	ostr := b.Original
	for i := range b.BBCodes {
		if strings.ToLower(b.BBCodes[i].Name) != "url" || b.BBCodes[i].IsClose || !b.BBCodes[i].IsValid {
			continue
		}
		nstr = cutString(nstr, b.BBCodes[i].Pos-1, b.BBCodes[i].Pos-1+b.BBCodes[i].Len)
		nstr = strings.Repeat(" ", b.BBCodes[i].Len) + nstr
		ostr = cutString(ostr, b.BBCodes[i].OriginalStart, b.BBCodes[b.BBCodes[i].OpenFor].OriginalEnd)
		ostr = strings.Repeat(" ", b.BBCodes[b.BBCodes[i].OpenFor].OriginalEnd-b.BBCodes[i].OriginalStart+1) + ostr
	}
	lns := urlre.FindAllIndex([]byte(nstr), -1)
	prm := urlre.FindAll([]byte([]byte(nstr)), -1)
	org := urlre.FindAllIndex([]byte(b.Original), -1)

	for i := range lns {
		if len(lns[i]) == 2 {
			bopen := BBCode{
				OriginalStart: org[i][0],
				OriginalEnd:   org[i][1] - org[i][0],
				Pos:           lns[i][0],
				Len:           lns[i][1] - lns[i][0],
				Param:         string(prm[i]),
				Name:          "url",
				IsClose:       false,
				IsValid:       true,
				CloseFor:      -1,
				OpenFor:       len(b.BBCodes) + 1,
			}
			bclose := BBCode{
				OriginalStart: org[i][1],
				OriginalEnd:   org[i][1] + 3, // len(url) = 3
				Pos:           lns[i][1],
				Len:           0,
				Name:          "url",
				IsClose:       true,
				IsValid:       true,
				CloseFor:      len(b.BBCodes),
				OpenFor:       -1,
			}
			b.BBCodes = append(b.BBCodes, bopen, bclose)
		}
	}
}

func getRunesRange(s string, start, end int) (string, int) {
	var ra []rune
	rs := []rune(s)
	for i := range rs {
		if i >= start && i < end {
			ra = append(ra, rs[i])
		}
	}
	return string(ra), UTF16Count(s)
}

func getRunesLen(s string) (l int) {
	return len([]rune(s))
}

// UTF16Count get len of string like javascript
func UTF16Count(s string) (c int) {
	rs := []rune(s)
	for i := range rs {
		if utf8.RuneLen(rs[i]) == 1 {
			c++
			continue
		}
		c += utf8.RuneLen(rs[i]) / 2
	}
	return c
}
