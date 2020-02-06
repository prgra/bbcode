# bbcode

[![Go Report Card](https://goreportcard.com/badge/github.com/prgra/bbcode)](https://goreportcard.com/report/github.com/prgra/bbcode)

### golang bbcode parser

usage:

```go
package main

import (
	"github.com/prgr/bbcode"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := bbcode.Parse("[b]hello[/b][code]test[/code]")
	spew.Dump(a)
}

```

result:

```go
(bbcode.BBCodes) {
 Original: (string) (len=29) "[b]hello[/b][code]test[/code]",
 NewString: (string) (len=9) "hellotest",
 BBCodes: ([]bbcode.BBCode) (len=4 cap=4) {
  (bbcode.BBCode) {
   OriginalStart: (int) 1,
   OriginalEnd: (int) 3,
   Pos: (int) 1,
   Len: (int) 5,
   Param: (string) "",
   Name: (string) (len=1) "b",
   IsClose: (bool) false,
   IsValid: (bool) true,
   CloseFor: (int) -1,
   OpenFor: (int) 1
  },
  (bbcode.BBCode) {
   OriginalStart: (int) 9,
   OriginalEnd: (int) 12,
   Pos: (int) 6,
   Len: (int) 0,
   Param: (string) "",
   Name: (string) (len=1) "b",
   IsClose: (bool) true,
   IsValid: (bool) true,
   CloseFor: (int) 0,
   OpenFor: (int) -1
  },
  (bbcode.BBCode) {
   OriginalStart: (int) 13,
   OriginalEnd: (int) 18,
   Pos: (int) 6,
   Len: (int) 4,
   Param: (string) "",
   Name: (string) (len=4) "code",
   IsClose: (bool) false,
   IsValid: (bool) true,
   CloseFor: (int) -1,
   OpenFor: (int) 3
  },
  (bbcode.BBCode) {
   OriginalStart: (int) 23,
   OriginalEnd: (int) 29,
   Pos: (int) 10,
   Len: (int) 0,
   Param: (string) "",
   Name: (string) (len=4) "code",
   IsClose: (bool) true,
   IsValid: (bool) true,
   CloseFor: (int) 2,
   OpenFor: (int) -1
  }
 }
}
```
