package tmpl

import (
	"bytes"
	"unicode/utf8"
)

// readCode takes a string such as:
//
//	".Apple}} more stuff"
//
// and returns:
//
//	return ".Apple", " more stuff", nil
func readCode(s string) (string, string, error) {

	var code bytes.Buffer

	index := 0
	//Loop: for 0 < len(s[index:]) {
	Loop: for index < len(s) {

		r, length := utf8.DecodeRuneInString(s[index:])
		if utf8.RuneError == r {
			return s, "", errRuneError
		}
		index += length

		switch r {
		case '}':
			if len(s) < index {
				return s, "", errInternalError
			}
			r2, length2 := utf8.DecodeRuneInString(s[index:])
			if utf8.RuneError == r {
				return s, "", errRuneError
			}
			index += length2

			switch r2 {
			case '}':
				break Loop
			default:
				code.WriteRune(r)
				code.WriteRune(r2)
			}
		default:
			code.WriteRune(r)
		}
	}
	if len(s) < index {
		return s, "", errInternalError
	}

	return s[index:], code.String(), nil
}
