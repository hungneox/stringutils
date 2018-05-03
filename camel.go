package stringutils

import (
	"strings"
	"unicode"
)

// CamelCase convert snake_case to camelCase
func CamelCase(value string) string {
	output := ""
	input := strings.TrimSpace(value)
	upperNext := false
	for _, r := range input {
		if r == '_' || r == '-' || r == ' ' {
			upperNext = true
		} else {
			if upperNext {
				output += strings.ToUpper(string(r))
			} else {
				output += string(r)
			}

			if unicode.IsNumber(r) {
				upperNext = true
			} else {
				upperNext = false
			}
		}
	}

	return output
}

// PascalCase convert snake_case to PascalCase
func PascalCase(value string) string {
	return strings.Title(CamelCase(value))
}
