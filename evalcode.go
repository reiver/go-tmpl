package tmpl

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// evalCode evaluates `code`, using the data in `data`, writing the output
// to `w`.
//
// Example
//
//	var buffer bytes.Buffer
//	
//	code := ".Name"
//	
//	data := map[string][]interface{}{
//		"Name": []interface{}{
//			"Joe Blow"
//		},
//		"Age": []interface{}{
//			41
//		}
//		"Sex": []interface{}{
//			"male",
//		}
//	}
//	
//	if err := evalCode(&buffer, name, data); nil != err {
//		return err
//	}
//	
//	io.WriteString(os.Stdout, buffer.String())
//	
//	// Output:
//	// Joe Blow
func evalCode(w io.Writer, code string, data map[string][]interface{}) error {

	r, length := utf8.DecodeRuneInString(code)
	switch r {
	case '.':
		fieldName := code[length:]

		if values, ok := data[fieldName]; !ok {
			writeError(w, fmt.Sprintf("Not Found: %q", code))
		} else if 1 > len(values) {
			writeError(w, fmt.Sprintf("Not Found: %q", code))
		} else {

			value := values[0]

			switch x := value.(type) {
			case fmt.Stringer:
				io.WriteString(w, x.String())
			case interface { String()(string, error) }:
				s, err := x.String()
				if nil != err {
					return err
				}

				io.WriteString(w, s)
			default:
				fmt.Fprintf(w, "%v", value)
			}
		}
	default:
		writeError(w, fmt.Sprintf("Bad Code: %q", code))
	}

	return nil
}

func writeError(w io.Writer, a ...string) {

	io.WriteString(w, "{{::ERROR")
	if 1 <= len(a) {
		io.WriteString(w, ":")

		for _, s := range a {
			io.WriteString(w, " ")
			io.WriteString(w, s)
		}
	}
	io.WriteString(w, "::}}")
}
