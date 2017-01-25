package tmpl

import (
	"bytes"

	"testing"
)

func TestFprintt(t *testing.T) {

	tests := []struct{
		Template string
		Data     interface{}
		Expected string
	}{
		{
			// Nothing here.
		},



		{
			Template: "{1} {{.Thing}} {, ,}",
			Data: struct{
				Thing bool
			}{
				Thing: true,
			},
			Expected: "{1} true {, ,}",
		},



		{
			Template: "(a) {{.Apple}} (b)",
			Data: struct{
				Apple bool
			}{
				Apple: true,
			},
			Expected: "(a) true (b)",
		},
		{
			Template: "(a) {{.Apple}} (b) {{.Banana}}, (c)",
			Data: struct{
				Apple  bool
				Banana float64
			}{
				Apple:  true,
				Banana: 2.2,
			},
			Expected: "(a) true (b) 2.2, (c)",
		},
		{
			Template: "(a) {{.Apple}} (b) {{.Banana}}, (c) {{.Cherry}} (d)",
			Data: struct{
				Apple  bool
				Banana float64
				Cherry int64
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
			},
			Expected: "(a) true (b) 2.2, (c) -3 (d)",
		},
		{
			Template: "(a) {{.Apple}} (b) {{.Banana}}, (c) {{.Cherry}} (d) {{.Date}} (e)",
			Data: struct{
				Apple  bool
				Banana float64
				Cherry int64
				Date   string
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
				Date:  "FOUR",
			},
			Expected: "(a) true (b) 2.2, (c) -3 (d) FOUR (e)",
		},



		{
			Template: "(a) {{.Apple}} (b) {{.Banana}}, (c) {{.Cherry}} (d) {{.Date}} (e) {{.Watermelon}} (f)",
			Data: struct{
				One        bool      `tmpl.name:"Apple"`
				Two        float64   `tmpl:"Banana"`
				Three      int64     `tmpl.name:"Cherry"`
				Four       string    `tmpl:"Date"`
				Watermelon complex64
				ShouldNotShow uint8
			}{
				One:       true,
				Two:        2.2,
				Three:      -3,
				Four:       "FOUR",
				Watermelon: -5i,
				ShouldNotShow: 6,
			},
			Expected: "(a) true (b) 2.2, (c) -3 (d) FOUR (e) (0-5i) (f)",
		},



		{
			Template: "{{",
			Data: struct{
				Thing bool
			}{
				Thing: true,
			},
			Expected: "{{::ERROR: Bad Code: \"\"::}}",
		},



		{
			Template: "{{.Thing}} {{",
			Data: struct{
				Thing bool
			}{
				Thing: true,
			},
			Expected: "true {{::ERROR: Bad Code: \"\"::}}",
		},



		{
			Template: "{{.Apple}} {{.Thing}} {{.Banana}}",
			Data: struct{
				Apple  bool
				Banana float64
				Cherry int64
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
			},
			Expected: "true {{::ERROR: Not Found: \".Thing\"::}} 2.2",
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		actualN, err := Fprintt(&buffer, test.Template, test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Expected), actualN; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("EXPECTED String: %q", test.Expected)
			t.Errorf("ACTUAL String:   %q", buffer.String())
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d,...", testNumber)
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			continue
		}
	}
}
