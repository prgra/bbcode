package bbcode

import "regexp"

var urlrestr = string([]byte{40, 63, 105, 41, 40, 40, 97, 97, 97, 124, 97, 97, 97, 115, 124, 97, 98, 111, 117, 116, 124, 97, 99, 97, 112, 124, 97, 99, 99, 116, 124, 97, 99, 100, 124, 97, 99, 114, 124, 97, 100, 105, 117, 109, 120, 116, 114, 97, 124, 97, 100, 116, 124, 97, 102, 112, 124, 97, 102, 115, 124, 97, 105, 109, 124, 97, 109, 115, 115, 124, 97, 110, 100, 114, 111, 105, 100, 124, 97, 112, 112, 100, 97, 116, 97, 124, 97, 112, 116, 124, 97, 114, 107, 124, 97, 116, 116, 97, 99, 104, 109, 101, 110, 116, 124, 97, 119, 124, 98, 97, 114, 105, 111, 110, 124, 98, 101, 115, 104, 97, 114, 101, 124, 98, 105, 116, 99, 111, 105, 110, 124, 98, 105, 116, 99, 111, 105, 110, 99, 97, 115, 104, 124, 98, 108, 111, 98, 124, 98, 111, 108, 111, 124, 98, 114, 111, 119, 115, 101, 114, 101, 120, 116, 124, 99, 97, 108, 99, 117, 108, 97, 116, 111, 114, 124, 99, 97, 108, 108, 116, 111, 124, 99, 97, 112, 124, 99, 97, 115, 116, 124, 99, 97, 115, 116, 115, 124, 99, 104, 114, 111, 109, 101, 124, 99, 104, 114, 111, 109, 101, 45, 101, 120, 116, 101, 110, 115, 105, 111, 110, 124, 99, 105, 100, 124, 99, 111, 97, 112, 124, 99, 111, 97, 112, 92, 43, 116, 99, 112, 124, 99, 111, 97, 112, 92, 43, 119, 115, 124, 99, 111, 97, 112, 115, 124, 99, 111, 97, 112, 115, 92, 43, 116, 99, 112, 124, 99, 111, 97, 112, 115, 92, 43, 119, 115, 124, 99, 111, 109, 45, 101, 118, 101, 110, 116, 98, 114, 105, 116, 101, 45, 97, 116, 116, 101, 110, 100, 101, 101, 124, 99, 111, 110, 116, 101, 110, 116, 124, 99, 111, 110, 116, 105, 124, 99, 114, 105, 100, 124, 99, 118, 115, 124, 100, 97, 98, 124, 100, 97, 116, 97, 124, 100, 97, 118, 124, 100, 105, 97, 115, 112, 111, 114, 97, 124, 100, 105, 99, 116, 124, 100, 105, 100, 124, 100, 105, 115, 124, 100, 108, 110, 97, 45, 112, 108, 97, 121, 99, 111, 110, 116, 97, 105, 110, 101, 114, 124, 100, 108, 110, 97, 45, 112, 108, 97, 121, 115, 105, 110, 103, 108, 101, 124, 100, 110, 115, 124, 100, 110, 116, 112, 124, 100, 112, 112, 124, 100, 114, 109, 124, 100, 114, 111, 112, 124, 100, 116, 109, 105, 124, 100, 116, 110, 124, 100, 118, 98, 124, 101, 100, 50, 107, 124, 101, 108, 115, 105, 124, 101, 120, 97, 109, 112, 108, 101, 124, 102, 97, 99, 101, 116, 105, 109, 101, 124, 102, 97, 120, 124, 102, 101, 101, 100, 124, 102, 101, 101, 100, 114, 101, 97, 100, 121, 124, 102, 105, 108, 101, 124, 102, 105, 108, 101, 115, 121, 115, 116, 101, 109, 124, 102, 105, 110, 103, 101, 114, 124, 102, 105, 114, 115, 116, 45, 114, 117, 110, 45, 112, 101, 110, 45, 101, 120, 112, 101, 114, 105, 101, 110, 99, 101, 124, 102, 105, 115, 104, 124, 102, 109, 124, 102, 116, 112, 124, 102, 117, 99, 104, 115, 105, 97, 45, 112, 107, 103, 124, 103, 101, 111, 124, 103, 103, 124, 103, 105, 116, 124, 103, 105, 122, 109, 111, 112, 114, 111, 106, 101, 99, 116, 124, 103, 111, 124, 103, 111, 112, 104, 101, 114, 124, 103, 114, 97, 112, 104, 124, 103, 116, 97, 108, 107, 124, 104, 51, 50, 51, 124, 104, 97, 109, 124, 104, 99, 97, 112, 124, 104, 99, 112, 124, 104, 116, 116, 112, 124, 104, 116, 116, 112, 115, 124, 104, 120, 120, 112, 124, 104, 120, 120, 112, 115, 124, 104, 121, 100, 114, 97, 122, 111, 110, 101, 124, 105, 97, 120, 124, 105, 99, 97, 112, 124, 105, 99, 111, 110, 124, 105, 109, 124, 105, 109, 97, 112, 124, 105, 110, 102, 111, 124, 105, 111, 116, 100, 105, 115, 99, 111, 124, 105, 112, 110, 124, 105, 112, 112, 124, 105, 112, 112, 115, 124, 105, 114, 99, 124, 105, 114, 99, 54, 124, 105, 114, 99, 115, 124, 105, 114, 105, 115, 124, 105, 114, 105, 115, 92, 46, 98, 101, 101, 112, 124, 105, 114, 105, 115, 92, 46, 108, 119, 122, 124, 105, 114, 105, 115, 92, 46, 120, 112, 99, 124, 105, 114, 105, 115, 92, 46, 120, 112, 99, 115, 124, 105, 115, 111, 115, 116, 111, 114, 101, 124, 105, 116, 109, 115, 124, 106, 97, 98, 98, 101, 114, 124, 106, 97, 114, 124, 106, 109, 115, 124, 107, 101, 121, 112, 97, 114, 99, 124, 108, 97, 115, 116, 102, 109, 124, 108, 100, 97, 112, 124, 108, 100, 97, 112, 115, 124, 108, 101, 97, 112, 116, 111, 102, 114, 111, 103, 97, 110, 115, 124, 108, 111, 114, 97, 119, 97, 110, 124, 108, 118, 108, 116, 124, 109, 97, 103, 110, 101, 116, 124, 109, 97, 105, 108, 115, 101, 114, 118, 101, 114, 124, 109, 97, 105, 108, 116, 111, 124, 109, 97, 112, 115, 124, 109, 97, 114, 107, 101, 116, 124, 109, 101, 115, 115, 97, 103, 101, 124, 109, 105, 99, 114, 111, 115, 111, 102, 116, 92, 46, 119, 105, 110, 100, 111, 119, 115, 92, 46, 99, 97, 109, 101, 114, 97, 124, 109, 105, 99, 114, 111, 115, 111, 102, 116, 92, 46, 119, 105, 110, 100, 111, 119, 115, 92, 46, 99, 97, 109, 101, 114, 97, 92, 46, 109, 117, 108, 116, 105, 112, 105, 99, 107, 101, 114, 124, 109, 105, 99, 114, 111, 115, 111, 102, 116, 92, 46, 119, 105, 110, 100, 111, 119, 115, 92, 46, 99, 97, 109, 101, 114, 97, 92, 46, 112, 105, 99, 107, 101, 114, 124, 109, 105, 100, 124, 109, 109, 115, 124, 109, 111, 100, 101, 109, 124, 109, 111, 110, 103, 111, 100, 98, 124, 109, 111, 122, 124, 109, 115, 45, 97, 99, 99, 101, 115, 115, 124, 109, 115, 45, 98, 114, 111, 119, 115, 101, 114, 45, 101, 120, 116, 101, 110, 115, 105, 111, 110, 124, 109, 115, 45, 99, 97, 108, 99, 117, 108, 97, 116, 111, 114, 124, 109, 115, 45, 100, 114, 105, 118, 101, 45, 116, 111, 124, 109, 115, 45, 101, 110, 114, 111, 108, 108, 109, 101, 110, 116, 124, 109, 115, 45, 101, 120, 99, 101, 108, 124, 109, 115, 45, 101, 121, 101, 99, 111, 110, 116, 114, 111, 108, 115, 112, 101, 101, 99, 104, 124, 109, 115, 45, 103, 97, 109, 101, 98, 97, 114, 115, 101, 114, 118, 105, 99, 101, 115, 124, 109, 115, 45, 103, 97, 109, 105, 110, 103, 111, 118, 101, 114, 108, 97, 121, 124, 109, 115, 45, 103, 101, 116, 111, 102, 102, 105, 99, 101, 124, 109, 115, 45, 104, 101, 108, 112, 124, 109, 115, 45, 105, 110, 102, 111, 112, 97, 116, 104, 124, 109, 115, 45, 105, 110, 112, 117, 116, 97, 112, 112, 124, 109, 115, 45, 108, 111, 99, 107, 115, 99, 114, 101, 101, 110, 99, 111, 109, 112, 111, 110, 101, 110, 116, 45, 99, 111, 110, 102, 105, 103, 124, 109, 115, 45, 109, 101, 100, 105, 97, 45, 115, 116, 114, 101, 97, 109, 45, 105, 100, 124, 109, 115, 45, 109, 105, 120, 101, 100, 114, 101, 97, 108, 105, 116, 121, 99, 97, 112, 116, 117, 114, 101, 124, 109, 115, 45, 109, 111, 98, 105, 108, 101, 112, 108, 97, 110, 115, 124, 109, 115, 45, 111, 102, 102, 105, 99, 101, 97, 112, 112, 124, 109, 115, 45, 112, 101, 111, 112, 108, 101, 124, 109, 115, 45, 112, 114, 111, 106, 101, 99, 116, 124, 109, 115, 45, 112, 111, 119, 101, 114, 112, 111, 105, 110, 116, 124, 109, 115, 45, 112, 117, 98, 108, 105, 115, 104, 101, 114, 124, 109, 115, 45, 114, 101, 115, 116, 111, 114, 101, 116, 97, 98, 99, 111, 109, 112, 97, 110, 105, 111, 110, 124, 109, 115, 45, 115, 99, 114, 101, 101, 110, 99, 108, 105, 112, 124, 109, 115, 45, 115, 99, 114, 101, 101, 110, 115, 107, 101, 116, 99, 104, 124, 109, 115, 45, 115, 101, 97, 114, 99, 104, 124, 109, 115, 45, 115, 101, 97, 114, 99, 104, 45, 114, 101, 112, 97, 105, 114, 124, 109, 115, 45, 115, 101, 99, 111, 110, 100, 97, 114, 121, 45, 115, 99, 114, 101, 101, 110, 45, 99, 111, 110, 116, 114, 111, 108, 108, 101, 114, 124, 109, 115, 45, 115, 101, 99, 111, 110, 100, 97, 114, 121, 45, 115, 99, 114, 101, 101, 110, 45, 115, 101, 116, 117, 112, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 97, 105, 114, 112, 108, 97, 110, 101, 109, 111, 100, 101, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 98, 108, 117, 101, 116, 111, 111, 116, 104, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 99, 97, 109, 101, 114, 97, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 99, 101, 108, 108, 117, 108, 97, 114, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 99, 108, 111, 117, 100, 115, 116, 111, 114, 97, 103, 101, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 99, 111, 110, 110, 101, 99, 116, 97, 98, 108, 101, 100, 101, 118, 105, 99, 101, 115, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 100, 105, 115, 112, 108, 97, 121, 115, 45, 116, 111, 112, 111, 108, 111, 103, 121, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 101, 109, 97, 105, 108, 97, 110, 100, 97, 99, 99, 111, 117, 110, 116, 115, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 108, 97, 110, 103, 117, 97, 103, 101, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 108, 111, 99, 97, 116, 105, 111, 110, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 108, 111, 99, 107, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 110, 102, 99, 116, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 110, 111, 116, 105, 102, 105, 99, 97, 116, 105, 111, 110, 115, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 112, 111, 119, 101, 114, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 112, 114, 105, 118, 97, 99, 121, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 112, 114, 111, 120, 105, 109, 105, 116, 121, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 115, 99, 114, 101, 101, 110, 114, 111, 116, 97, 116, 105, 111, 110, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 119, 105, 102, 105, 124, 109, 115, 45, 115, 101, 116, 116, 105, 110, 103, 115, 45, 119, 111, 114, 107, 112, 108, 97, 99, 101, 124, 109, 115, 45, 115, 112, 100, 124, 109, 115, 45, 115, 116, 116, 111, 118, 101, 114, 108, 97, 121, 124, 109, 115, 45, 116, 114, 97, 110, 115, 105, 116, 45, 116, 111, 124, 109, 115, 45, 117, 115, 101, 114, 97, 99, 116, 105, 118, 105, 116, 121, 115, 101, 116, 124, 109, 115, 45, 118, 105, 114, 116, 117, 97, 108, 116, 111, 117, 99, 104, 112, 97, 100, 124, 109, 115, 45, 118, 105, 115, 105, 111, 124, 109, 115, 45, 119, 97, 108, 107, 45, 116, 111, 124, 109, 115, 45, 119, 104, 105, 116, 101, 98, 111, 97, 114, 100, 124, 109, 115, 45, 119, 104, 105, 116, 101, 98, 111, 97, 114, 100, 45, 99, 109, 100, 124, 109, 115, 45, 119, 111, 114, 100, 124, 109, 115, 110, 105, 109, 124, 109, 115, 114, 112, 124, 109, 115, 114, 112, 115, 124, 109, 115, 115, 124, 109, 116, 113, 112, 124, 109, 117, 109, 98, 108, 101, 124, 109, 117, 112, 100, 97, 116, 101, 124, 109, 118, 110, 124, 110, 101, 119, 115, 124, 110, 102, 115, 124, 110, 105, 124, 110, 105, 104, 124, 110, 110, 116, 112, 124, 110, 111, 116, 101, 115, 124, 111, 99, 102, 124, 111, 105, 100, 124, 111, 110, 101, 110, 111, 116, 101, 124, 111, 110, 101, 110, 111, 116, 101, 45, 99, 109, 100, 124, 111, 112, 97, 113, 117, 101, 108, 111, 99, 107, 116, 111, 107, 101, 110, 124, 111, 112, 101, 110, 112, 103, 112, 52, 102, 112, 114, 124, 112, 97, 99, 107, 124, 112, 97, 108, 109, 124, 112, 97, 112, 97, 114, 97, 122, 122, 105, 124, 112, 97, 121, 109, 101, 110, 116, 124, 112, 97, 121, 116, 111, 124, 112, 107, 99, 115, 49, 49, 124, 112, 108, 97, 116, 102, 111, 114, 109, 124, 112, 111, 112, 124, 112, 114, 101, 115, 124, 112, 114, 111, 115, 112, 101, 114, 111, 124, 112, 114, 111, 120, 121, 124, 112, 119, 105, 100, 124, 112, 115, 121, 99, 124, 112, 116, 116, 112, 124, 113, 98, 124, 113, 117, 101, 114, 121, 124, 113, 117, 105, 99, 45, 116, 114, 97, 110, 115, 112, 111, 114, 116, 124, 114, 101, 100, 105, 115, 124, 114, 101, 100, 105, 115, 115, 124, 114, 101, 108, 111, 97, 100, 124, 114, 101, 115, 124, 114, 101, 115, 111, 117, 114, 99, 101, 124, 114, 109, 105, 124, 114, 115, 121, 110, 99, 124, 114, 116, 109, 102, 112, 124, 114, 116, 109, 112, 124, 114, 116, 115, 112, 124, 114, 116, 115, 112, 115, 124, 114, 116, 115, 112, 117, 124, 115, 101, 99, 111, 110, 100, 108, 105, 102, 101, 124, 115, 101, 114, 118, 105, 99, 101, 124, 115, 101, 115, 115, 105, 111, 110, 124, 115, 102, 116, 112, 124, 115, 103, 110, 124, 115, 104, 116, 116, 112, 124, 115, 105, 101, 118, 101, 124, 115, 105, 109, 112, 108, 101, 108, 101, 100, 103, 101, 114, 124, 115, 105, 112, 124, 115, 105, 112, 115, 124, 115, 107, 121, 112, 101, 124, 115, 109, 98, 124, 115, 109, 115, 124, 115, 109, 116, 112, 124, 115, 110, 101, 119, 115, 124, 115, 110, 109, 112, 124, 115, 111, 97, 112, 92, 46, 98, 101, 101, 112, 124, 115, 111, 97, 112, 92, 46, 98, 101, 101, 112, 115, 124, 115, 111, 108, 100, 97, 116, 124, 115, 112, 105, 102, 102, 101, 124, 115, 112, 111, 116, 105, 102, 121, 124, 115, 115, 104, 124, 115, 116, 101, 97, 109, 124, 115, 116, 117, 110, 124, 115, 116, 117, 110, 115, 124, 115, 117, 98, 109, 105, 116, 124, 115, 118, 110, 124, 116, 97, 103, 124, 116, 101, 97, 109, 115, 112, 101, 97, 107, 124, 116, 101, 108, 124, 116, 101, 108, 105, 97, 101, 105, 100, 124, 116, 101, 108, 110, 101, 116, 124, 116, 102, 116, 112, 124, 116, 104, 105, 110, 103, 115, 124, 116, 104, 105, 115, 109, 101, 115, 115, 97, 103, 101, 124, 116, 105, 112, 124, 116, 110, 51, 50, 55, 48, 124, 116, 111, 111, 108, 124, 116, 117, 114, 110, 124, 116, 117, 114, 110, 115, 124, 116, 118, 124, 117, 100, 112, 124, 117, 110, 114, 101, 97, 108, 124, 117, 114, 110, 124, 117, 116, 50, 48, 48, 52, 124, 118, 45, 101, 118, 101, 110, 116, 124, 118, 101, 109, 109, 105, 124, 118, 101, 110, 116, 114, 105, 108, 111, 124, 118, 105, 100, 101, 111, 116, 101, 120, 124, 118, 110, 99, 124, 118, 105, 101, 119, 45, 115, 111, 117, 114, 99, 101, 124, 119, 97, 105, 115, 124, 119, 101, 98, 99, 97, 108, 124, 119, 112, 105, 100, 124, 119, 115, 124, 119, 115, 115, 124, 119, 116, 97, 105, 124, 119, 121, 99, 105, 119, 121, 103, 124, 120, 99, 111, 110, 124, 120, 99, 111, 110, 45, 117, 115, 101, 114, 105, 100, 124, 120, 102, 105, 114, 101, 124, 120, 109, 108, 114, 112, 99, 92, 46, 98, 101, 101, 112, 124, 120, 109, 108, 114, 112, 99, 92, 46, 98, 101, 101, 112, 115, 124, 120, 109, 112, 112, 124, 120, 114, 105, 124, 121, 109, 115, 103, 114, 124, 122, 51, 57, 92, 46, 53, 48, 124, 122, 51, 57, 92, 46, 53, 48, 114, 124, 122, 51, 57, 92, 46, 53, 48, 115, 41, 58, 47, 47, 124, 40, 98, 105, 116, 99, 111, 105, 110, 124, 102, 105, 108, 101, 124, 109, 97, 103, 110, 101, 116, 124, 109, 97, 105, 108, 116, 111, 124, 115, 109, 115, 124, 116, 101, 108, 124, 120, 109, 112, 112, 41, 58, 41, 40, 63, 45, 105, 41, 40, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 40, 92, 40, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 40, 92, 40, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 92, 41, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 41, 42, 92, 41, 124, 92, 91, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 40, 92, 91, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 92, 93, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 41, 42, 92, 93, 124, 92, 123, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 40, 92, 123, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 92, 125, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 95, 42, 92, 112, 123, 80, 111, 125, 93, 42, 41, 42, 92, 125, 124, 91, 92, 112, 123, 76, 125, 92, 112, 123, 77, 125, 92, 112, 123, 78, 125, 47, 92, 45, 95, 43, 38, 126, 37, 61, 35, 92, 112, 123, 83, 99, 125, 92, 112, 123, 83, 111, 125, 93, 41, 43, 41, 43})

var urlre *regexp.Regexp

func init() {

	urlre = regexp.MustCompile(urlrestr)
}