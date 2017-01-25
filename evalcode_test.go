package tmpl

import (
	"bytes"

	"testing"
)

func TestEvalCode(t *testing.T) {

	tests := []struct{
		Code string
		Data map[string][]interface{}
		Expected string
	}{
		{
			Code: "",
			Data: map[string][]interface{}{},
			Expected: "{{::ERROR: Bad Code: \"\"::}}",
		},
		{
			Code: ".Thing",
			Data: map[string][]interface{}{},
			Expected: "{{::ERROR: Not Found: \".Thing\"::}}",
		},
		{
			Code: "Thing",
			Data: map[string][]interface{}{},
			Expected: "{{::ERROR: Bad Code: \"Thing\"::}}",
		},



		{
			Code: "",
			Data: map[string][]interface{}{
				"Thing": []interface{}{
					int64(5),
				},
			},
			Expected: "{{::ERROR: Bad Code: \"\"::}}",
		},
		{
			Code: ".Thing",
			Data: map[string][]interface{}{
				"Thing": []interface{}{
					int64(5),
				},
			},
			Expected: "5",
		},
		{
			Code: "Thing",
			Data: map[string][]interface{}{
				"Thing": []interface{}{
					int64(5),
				},
			},
			Expected: "{{::ERROR: Bad Code: \"Thing\"::}}",
		},



		{
			Code: "",
			Data: map[string][]interface{}{
				"Other": []interface{}{
					int64(5),
				},
			},
			Expected: "{{::ERROR: Bad Code: \"\"::}}",
		},
		{
			Code: ".Thing",
			Data: map[string][]interface{}{
				"Other": []interface{}{
					int64(5),
				},
			},
			Expected: "{{::ERROR: Not Found: \".Thing\"::}}",
		},
		{
			Code: "Thing",
			Data: map[string][]interface{}{
				"Other": []interface{}{
					int64(5),
				},
			},
			Expected: "{{::ERROR: Bad Code: \"Thing\"::}}",
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if err := evalCode(&buffer, test.Code, test.Data); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
