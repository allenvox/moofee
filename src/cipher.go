package main

import (
	"strings"
)

func caesar(text string, shift int) string {
	chars := []rune(strings.ToLower(text))
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 'a' && chars[i] <= 'z' { // english
			shift = (shift%26 + 26) % 26 // only between 0 and 25
		} else if chars[i] >= 'а' && chars[i] <= 'я' { // russian
			shift = (shift%33 + 33) % 33 // only between 0 and 32
		} else {
			continue // skip
		}
		char := chars[i] + rune(shift)
		if chars[i] >= 'a' && chars[i] <= 'z' {
			if char > 'z' {
				char -= 26
			}
		} else if chars[i] >= 'а' && chars[i] <= 'я' {
			if char > 'я' {
				char -= 33
			}
		}
		chars[i] = char
	}
	return string(chars)
}

func vigenereEncode(text string, key string) string {
	textChars := []rune(strings.ToLower(text))
	keyChars := []rune(key)
	keylen := len(keyChars)
	letter := 0
	for i := 0; i < len(textChars); i++ {
		letter %= keylen
		char := textChars[i] + keyChars[letter]
		if textChars[i] >= 'a' && textChars[i] <= 'z' {
			if char > 'z' {
				char -= 'z'
			}
		} else if textChars[i] >= 'а' && textChars[i] <= 'я' {
			if char > 'я' {
				char -= 'я'
			}
		}
		textChars[i] = char
		letter++
	}
	return string(textChars)
}

func vigenereDecode(text string, key string) string {
	textChars := []rune(strings.ToLower(text))
	keyChars := []rune(key)
	keylen := len(keyChars)
	letter := 0
	for i := 0; i < len(textChars); i++ {
		letter %= keylen
		char := textChars[i] - keyChars[letter]
		if textChars[i] >= 'a' && textChars[i] <= 'z' {
			if char < 'a' {
				char += 'a'
			}
		} else if textChars[i] >= 'а' && textChars[i] <= 'я' {
			if char < 'а' {
				char += 'а'
			}
		}
		textChars[i] = char
		letter++
	}
	vigenere_switch = 0
	return string(textChars)
}
