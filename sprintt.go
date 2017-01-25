package tmpl

import (
	"bytes"
)

// Sprintt renders according to a template bound to the data and returns the resulting string.
func Sprintt(template string, data interface{}) string {
	var buffer bytes.Buffer

	Fprintt(&buffer, template, data)

	return buffer.String()
}
