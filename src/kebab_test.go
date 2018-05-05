package stringutils

import "testing"

func TestKebabCase(t *testing.T) {
	cases := [][]string{
		[]string{"foo_bar", "foo-bar"},
		[]string{"fooBar", "foo-bar"},
		[]string{"foo bar", "foo-bar"},
		[]string{" Foo bar baz", "foo-bar-baz"},
		[]string{"foo_bar_baz", "foo-bar-baz"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := KebabCase(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
