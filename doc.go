/*
Package tmpl provides templating capabilities.

The tmpl.Fprintt(), tmpl.Printt(), and tmpl.Sprintt() funcs are similar to
the Go built-in fmt.Fprintf(), fmt.Printf(), fmt.Sprintf();
except that it uses template-style for the format string.

For example:

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
*/
package tmpl
