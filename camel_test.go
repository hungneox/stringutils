package stringutils

import (
	"testing"
)

func TestPascalCase(t *testing.T) {
	cases := [][]string{
		[]string{"test_case", "TestCase"},
		[]string{"test", "Test"},
		[]string{"TestCase", "TestCase"},
		[]string{" test  case ", "TestCase"},
		[]string{"", ""},
		[]string{"many_many_words", "ManyManyWords"},
		[]string{"AnyKind of_string", "AnyKindOfString"},
		[]string{"odd-fix", "OddFix"},
		[]string{"numbers2And55with000", "Numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := PascalCase(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestCameCase(t *testing.T) {
	cases := [][]string{
		[]string{"foo-bar", "fooBar"},
		[]string{"foo_bar", "fooBar"},
		[]string{"foo  bar", "fooBar"},
		[]string{"foo_bar_baz", "fooBarBaz"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := CamelCase(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
