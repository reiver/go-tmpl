package tmpl

import (
	"os"
)

// Printt renders according to a template bound to the data and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printt(template string, data interface{}) (n int, err error) {
	return Fprintt(os.Stdout, template, data)
}
