package tmpl

import (
	"testing"
)

func TestReadCode(t *testing.T) {

	tests := []struct{
		Text     string
		ExpectedString string
		ExpectedCode   string
	}{
		{
			Text:           "apple}} banana cherry",
			ExpectedString:        " banana cherry",
			ExpectedCode:   "apple",
		},
		{
			Text:           "Apple}}Banana Cherry Date ",
			ExpectedString:        "Banana Cherry Date ",
			ExpectedCode:   "Apple",
		},



		{
			Text:           "something}}",
			ExpectedString:            "",
			ExpectedCode:   "something",
		},
		{
			Text:           "something}} ",
			ExpectedString:            " ",
			ExpectedCode:   "something",
		},



		{
			Text:           ".Apple}} {{.Banana}} {{.Cherry}}",
			ExpectedString:         " {{.Banana}} {{.Cherry}}",
			ExpectedCode:   ".Apple",
		},
		{
			Text:           ".Apple}}{{.Banana}}{{.Cherry}}",
			ExpectedString:         "{{.Banana}}{{.Cherry}}",
			ExpectedCode:   ".Apple",
		},



		{
			Text:           ".Apple.Banana.Cherry}}",
			ExpectedString:                       "",
			ExpectedCode:   ".Apple.Banana.Cherry",
		},
		{
			Text:           ".Apple.Banana.Cherry}}{{.Date}}",
			ExpectedString:                       "{{.Date}}",
			ExpectedCode:   ".Apple.Banana.Cherry",
		},
		{
			Text:           ".Apple.Banana.Cherry}} {{.Date}} ",
			ExpectedString:                       " {{.Date}} ",
			ExpectedCode:   ".Apple.Banana.Cherry",
		},



		{
			Text:           "}}",
			ExpectedString:    "",
			ExpectedCode:   "",
		},
		{
			Text:           "}}a",
			ExpectedString:   "a",
			ExpectedCode:   "",
		},
		{
			Text:           "}}{{a}}",
			ExpectedString:   "{{a}}",
			ExpectedCode:   "",
		},



		{
			Text:           " }}",
			ExpectedString:    "",
			ExpectedCode:   " ",
		},
		{
			Text:           " }}a",
			ExpectedString:    "a",
			ExpectedCode:   " ",
		},
		{
			Text:           " }}{{a}}",
			ExpectedString:    "{{a}}",
			ExpectedCode:   " ",
		},



		{
			Text:           ".}}",
			ExpectedString:    "",
			ExpectedCode:   ".",
		},
		{
			Text:           ".}}a",
			ExpectedString:    "a",
			ExpectedCode:   ".",
		},
		{
			Text:           ".}}{{a}}",
			ExpectedString:    "{{a}}",
			ExpectedCode:   ".",
		},



		{
			Text:           "e}}",
			ExpectedString:    "",
			ExpectedCode:   "e",
		},
		{
			Text:           "e}}a",
			ExpectedString:    "a",
			ExpectedCode:   "e",
		},
		{
			Text:           "e}}{{a}}",
			ExpectedString:    "{{a}}",
			ExpectedCode:   "e",
		},



		{
			Text:           "eF}}",
			ExpectedString:    "",
			ExpectedCode:   "eF",
		},
		{
			Text:           "eF}}a",
			ExpectedString:    "a",
			ExpectedCode:   "eF",
		},
		{
			Text:           "eF}}{{a}}",
			ExpectedString:    "{{a}}",
			ExpectedCode:   "eF",
		},



		{
			Text:           "🙂}}",
			ExpectedString:    "",
			ExpectedCode:   "🙂",
		},
		{
			Text:           "🙂}}a",
			ExpectedString:    "a",
			ExpectedCode:   "🙂",
		},
		{
			Text:           "🙂}}{{a}}",
			ExpectedString:    "{{a}}",
			ExpectedCode:   "🙂",
		},



		{
			Text:           "e}}",
			ExpectedString:    "",
			ExpectedCode:   "e",
		},
		{
			Text:           "e}}🙂",
			ExpectedString:    "🙂",
			ExpectedCode:   "e",
		},
		{
			Text:           "e}}{{🙂}}",
			ExpectedString:    "{{🙂}}",
			ExpectedCode:   "e",
		},



		{
			Text:           "hello world! 😁 🙂}}",
			ExpectedString:                     "",
			ExpectedCode:   "hello world! 😁 🙂",
		},
		{
			Text:           "hello world! 😁 🙂}}wow",
			ExpectedString:                     "wow",
			ExpectedCode:   "hello world! 😁 🙂",
		},
		{
			Text:           "hello world! 😁 🙂}}{{wow}}",
			ExpectedString:                     "{{wow}}",
			ExpectedCode:   "hello world! 😁 🙂",
		},
	}


	for testNumber, test := range tests {

		actualString, actualCode, err := readCode(test.Text)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.ExpectedString, actualString; expected != actual {
			t.Errorf("For test #%d,...", testNumber)
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			t.Errorf("EXPECTED CODE: %q", test.ExpectedCode)
			t.Errorf("ACTUAL CODE:   %q", actualCode)
			t.Errorf("TEXT: %q", test.Text)
			t.Errorf("")
			continue
		}

		if expected, actual := test.ExpectedCode, actualCode; expected != actual {
			t.Errorf("For test #%d,...", testNumber)
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			t.Errorf("TEXT: %q", test.Text)
			t.Errorf("")
			continue
		}
	}
}
