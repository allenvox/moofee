package main

import (
	"strings"
	"unicode"
)

func caesar(text string, shift int) string {
	shift = (shift%33 + 33) % 33 // only between 0 and 32
	chars := []rune(strings.ToLower(text))
	for i := 0; i < len(chars); i++ {
		if !unicode.IsLetter(chars[i]) {
			continue // skip
		}
		char := chars[i] + rune(shift)
		if char > 'я' {
			char -= 'я'
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
		if !unicode.IsLetter(textChars[i]) {
			continue
		}
		letter %= keylen
		char := textChars[i] + keyChars[letter]
		if char > 'я' {
			char -= 'я'
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
		if !unicode.IsLetter(textChars[i]) {
			continue
		}
		letter %= keylen
		char := textChars[i] - keyChars[letter]
		if char < 'а' {
			char += 'а'
		}
		textChars[i] = char
		letter++
	}
	vigenere_switch = 0
	return string(textChars)
}
