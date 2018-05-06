package stringutils

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "test_case"},
		[]string{"TestCase", "test_case"},
		[]string{"Test Case", "test_case"},
		[]string{" Test Case", "test_case"},
		[]string{"Test Case ", "test_case"},
		[]string{" Test Case ", "test_case"},
		[]string{"test-case", "test_case"},
		[]string{"many-many-words", "many_many_words"},
		[]string{"test", "test"},
		[]string{"test_case", "test_case"},
		[]string{"Test", "test"},
		[]string{"", ""},
		[]string{"ManyManyWords", "many_many_words"},
		[]string{"manyManyWords", "many_many_words"},
		[]string{"AnyKind of_string", "any_kind_of_string"},
		[]string{"numbers2and55with000", "numbers_2_and_55_with_000"},
		[]string{"JSONData", "json_data"},
	}
	for _, test := range cases {
		in := test[0]
		out := test[1]
		result := SnakeCase(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
