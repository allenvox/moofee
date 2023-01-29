package main

import (
	"strings"
)

func isRussian(letter rune) bool { // check if letter is russian
	if letter >= 'а' && letter <= 'я' {
		return true
	}
	return false
}

func isEnglish(letter rune) bool { // check if letter is english
	if letter >= 'a' && letter <= 'z' {
		return true
	}
	return false
}

func caesar(text string, shift int) string { // caesar cipher
	chars := []rune(strings.ToLower(text)) // get an array of lowercased characters
	for i := 0; i < len(chars); i++ {      // loop through characters
		if isEnglish(chars[i]) { // if letter is english
			shift = (shift%26 + 26) % 26 // only between 0 and 25
		} else if isRussian(chars[i]) { // russian
			shift = (shift%33 + 33) % 33 // only between 0 and 32
		} else {
			continue // skip if not a letter of any alphabets
		}
		char := chars[i] + rune(shift) // new char is an old char + shift
		if isEnglish(chars[i]) {
			if char > 'z' {
				char -= 26
			}
		} else if isRussian(chars[i]) {
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
		if isEnglish(textChars[i]) {
			if char > 'z' {
				char -= 'z'
			}
		} else if isRussian(textChars[i]) {
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
		if isEnglish(textChars[i]) {
			if char < 'a' {
				char += 'a'
			}
		} else if isRussian(textChars[i]) {
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
