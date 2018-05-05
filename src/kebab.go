package stringutils

import (
	"strings"
	"unicode"
)

// KebabCase convert snake_case to camelCase
func KebabCase(value string) string {
	output := ""
	input := strings.TrimSpace(value)

	for index, r := range input {
		if r == '_' || r == ' ' {
			output += "-"
		} else if unicode.IsUpper(r) && index > 0 {
			output += "-" + string(r)
		} else {
			output += string(r)
		}
	}

	return strings.ToLower(output)
}
