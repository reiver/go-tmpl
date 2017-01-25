package tmpl

import (
	"reflect"

	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct{
		Data interface{}
		Expected map[string][]interface{}
	}{
		{
			Data: struct{
				Apple  bool
			}{
				Apple:  true,
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true},
			},
		},
		{
			Data: struct{
				Apple  bool
				Banana float64
			}{
				Apple:  true,
				Banana: 2.2,
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true},
				"Banana": []interface{}{2.2},
			},
		},
		{
			Data: struct{
				Apple  bool
				Banana float64
				Cherry int64
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true},
				"Banana": []interface{}{2.2},
				"Cherry": []interface{}{int64(-3)},
			},
		},
		{
			Data: struct{
				Apple  bool
				Banana float64
				Cherry int64
				Date   string
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
				Date:   "FOUR",
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true},
				"Banana": []interface{}{2.2},
				"Cherry": []interface{}{int64(-3)},
				"Date":   []interface{}{"FOUR"},
			},
		},



		{
			Data: struct{
				One   bool    `tmpl.name:"Apple"`
				Two   float64 `tmpl:"Banana"`
				Three int64   `tmpl.name:"Cherry"`
				Four  string  `tmpl:"Date"`
				Watermelon complex128
			}{
				One:   true,
				Two:   2.2,
				Three: -3,
				Four:  "FOUR",
				Watermelon: -5i,
			},
			Expected: map[string][]interface{}{
				"Apple":      []interface{}{true},
				"Banana":     []interface{}{2.2},
				"Cherry":     []interface{}{int64(-3)},
				"Date":       []interface{}{"FOUR"},
				"Watermelon": []interface{}{-5i},
			},
		},



		{
			Data: struct{
				Apple  bool
				Banana float64 `tmpl.name:"Apple"`
				Cherry int64   `tmpl:"Apple"`
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true, float64(2.2), int64(-3)},
			},
		},
		{
			Data: struct{
				Apple  bool
				Banana float64 `tmpl.name:"Apple"`
				Cherry int64   `tmpl:"Apple"`
				Date   string
			}{
				Apple:  true,
				Banana: 2.2,
				Cherry: -3,
				Date:   "FOUR",
			},
			Expected: map[string][]interface{}{
				"Apple":  []interface{}{true, float64(2.2), int64(-3)},
				"Date":   []interface{}{"FOUR"},
			},
		},
	}

	for testNumber, test := range tests {

		actualData, err := extract(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Expected), len(actualData); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for key, expectedDatum := range test.Expected {
			if _, ok := actualData[key]; !ok {
				t.Errorf("For test #%d and key %q, expected key to be there, but actually wasn't.", testNumber, key)
				continue
			}

			if expected, actual := expectedDatum, actualData[key]; !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d and key %q, expected (%T) %v, but actually got (%T) %v.", testNumber, key, expected, expected, actual, actual)
				continue
			}
		}
	}
}
