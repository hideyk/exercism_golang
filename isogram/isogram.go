// Package isogram deals with anything isogram related
package isogram

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

func uniqueLetters(input string) bool {
	m := map[rune]bool{}
	for _, char := range input {
		_, ok := m[char]
		if ok {
			return false
		}
		m[char] = true
	}
	return true
}

// IsIsogram helps to check if a particular string is an isogram or not
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	reg, err := regexp.Compile("[^a-z]")
	if err != nil {
		os.Exit(100)
	}
	formattedInput := reg.ReplaceAllString(input, "")
	return uniqueLetters(formattedInput)
}

// IsIsogram2 is a more efficient and intuitive solution
func IsIsogram2(input string) bool {
	input = strings.ToLower(input)
	for i, char := range input {
		if unicode.IsLetter(char) && strings.ContainsRune(input[i+1:], char) {
			return false
		}
	}
	return true
}