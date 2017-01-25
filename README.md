# go-tmpl

Package tmpl provides templating capabilities, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-tmpl

[![GoDoc](https://godoc.org/github.com/reiver/go-tmpl?status.svg)](https://godoc.org/github.com/reiver/go-tmpl)


## Example
```go
import (
	"github.com/reiver/go-tmpl"
)

// ...

type Person struct {
	Name string
	Age  int
	Sex  string
}

var person Person

person.Name = "Joe Blow"
person.Age  = "41"
person.Sex  = "male"

tmpl.Printt("Hello {{.Name}}, you are {{.Age}} years old.", person)

// Output:
// Hello Joe Blow, you are 41 years old.
```
