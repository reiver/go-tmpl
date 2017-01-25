package tmpl

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// Fprintt renders according to a template bound to the data and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintt(w io.Writer, template string, data interface{}) (n int, err error) {

	// Turn whatever `data` is (probably a struct) into a
	// map[string][]interface{}.
	m, err := extract(data)
	if nil != err {
		return 0, err
	}
	if nil == m {
		m = map[string][]interface{}{}
	}

	var buffer bytes.Buffer

	s := template
	for 0 < len(s) {
		r, length := utf8.DecodeRuneInString(s)
		if utf8.RuneError == r {
			return 0, fmt.Errorf("Rune Error")
		}
		s = s[length:]

		switch r {
		case '{':

			r2, length2 := utf8.DecodeRuneInString(s)
			if utf8.RuneError == r {
				return 0, fmt.Errorf("Rune Error")
			}
			s = s[length2:]

			switch r2 {
			case '{':
				s2, code, err := readCode(s)
				if nil != err {
					return 0, err
				}
				s = s2

				if err := evalCode(&buffer, code, m); nil != err {
					return 0, err
				}
			default:
				buffer.WriteRune(r)
				buffer.WriteRune(r2)
			}
		default:
			buffer.WriteRune(r)
		}
	}

	return w.Write(buffer.Bytes())
}
