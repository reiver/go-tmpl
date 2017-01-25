package tmpl

import (
	"os"
)

func Printt(template string, data interface{}) (n int, err error) {
	return Fprintt(os.Stdout, template, data)
}
