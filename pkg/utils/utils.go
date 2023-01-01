package utils

import (
	"strings"
	"unicode"
)

func Split(s string, separators []rune) []string {
	f := func(r rune) bool {
		for _, s := range separators {
			if r == s {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(s, f)
}

func SanitizeAsMoney(input string) string {
	//allowed format is : [0-9]+[\.,]?[0-9]{0,2}
	//allowed characters are : [0-9\.,]
	sb := strings.Builder{}
	shouldStopOnNextComa := false
	digitsAfterComa := 0
	for _, r := range input {
		if digitsAfterComa == 2 {
			return sb.String()
		}
		if unicode.IsDigit(r) {
			if shouldStopOnNextComa {
				digitsAfterComa++
			}
			sb.WriteRune(r)
			continue
		}
		if r == ',' || r == '.' {
			if shouldStopOnNextComa {
				return sb.String()
			}
			sb.WriteRune(r)
			shouldStopOnNextComa = true
		}
	}
	return sb.String()
}
