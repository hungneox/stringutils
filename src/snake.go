package stringutils

import (
	"strings"
	"unicode"
)

// SnakeCase convert camelCase to snake_case
func SnakeCase(value string) string {
	output := ""
	// trim all spaces
	input := strings.TrimSpace(value)

	for index, r := range input {
		nextIsUpper := false
		nextIsNumber := false
		if index+1 < len(input) {
			nextIsUpper = unicode.IsUpper(rune(input[index+1]))
			nextIsNumber = unicode.IsNumber(rune(input[index+1]))
		}

		// treat acronym is a word, e.g: JSONData -> json_data
		if unicode.IsUpper(r) && !nextIsUpper && index > 0 {
			output += "_" + string(r)
		} else if r == ' ' || r == '-' {
			// Prevent add double "_"
			if !nextIsUpper {
				output += "_"
			}
		} else {
			output += string(r)
			if unicode.IsLetter(r) && nextIsNumber {
				output += "_"
			}
			if unicode.IsNumber(r) && !nextIsNumber && index != len(input)-1 {
				output += "_"
			}
		}
	}

	return strings.ToLower(output)
}
