package tmpl

import (
	"bytes"
)

func Sprintt(template string, data interface{}) string {
	var buffer bytes.Buffer

	Fprintt(&buffer, template, data)

	return buffer.String()
}
