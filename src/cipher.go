package main

import (
	"strings"
	"unicode"
)

func caesar(text string, shift int) string {
	shift = (shift%33 + 33) % 33 // only between 0 and 32
	text = strings.ToLower(text)
	chars := []rune(text)
	for i := 0; i < len(chars); i++ {
		if !unicode.IsLetter(chars[i]) {
			continue // skip
		}
		char := chars[i] + rune(shift)
		if char > 'я' {
			char = char - 'я' + 'а'
		}
		chars[i] = char
	}
	return string(chars)
}
