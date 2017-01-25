package tmpl

import (
	"errors"
)

var (
	errInternalError    = errors.New("Internal Error")
	errNilReflectedType = errors.New("Nil Reflected Type")
	errRuneError        = errors.New("Rune Error")
)
