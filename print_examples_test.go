package tmpl_test

import (
	"github.com/reiver/go-tmpl"

	"io"
	"os"
)

func ExampleFprintt() {

	type Person struct {
		Name string
		Age  int
		Sex  string
	}

	var person Person

	person.Name = "Joe Blow"
	person.Age  = 41
	person.Sex  = "male"

	tmpl.Fprintt(os.Stdout, "Hello {{.Name}}, you are {{.Age}} years old.", person)

	// Output:
	// Hello Joe Blow, you are 41 years old.
}

func ExamplePrintt() {

	type Person struct {
		Name string
		Age  int
		Sex  string
	}

	var person Person

	person.Name = "Joe Blow"
	person.Age  = 41
	person.Sex  = "male"

	tmpl.Printt("Hello {{.Name}}, you are {{.Age}} years old.", person)

	// Output:
	// Hello Joe Blow, you are 41 years old.
}

func ExampleSprintt() {

	type Person struct {
		Name string
		Age  int
		Sex  string
	}

	var person Person

	person.Name = "Joe Blow"
	person.Age  = 41
	person.Sex  = "male"

	s := tmpl.Sprintt("Hello {{.Name}}, you are {{.Age}} years old.", person)

	io.WriteString(os.Stdout, s)

	// Output:
	// Hello Joe Blow, you are 41 years old.
}

func ExampleSprintt_stringArray() {

	type Party struct {
		People []string
	}

	var party Party

	party.People = []string{"Joe Blow", "Douglas Wayne"}

	s := tmpl.Sprintt("The people attending the party are: {{.People}}", party)

	io.WriteString(os.Stdout, s)

	// Output:
	// The people attending the party are: [Joe Blow Douglas Wayne]
}
