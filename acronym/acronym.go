// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	reg, _ := regexp.Compile(`[^a-zA-Z\s]+`)
	s = strings.ReplaceAll(s, "-", " ")
	s = reg.ReplaceAllString(s, "")
	words := strings.Fields(s)
	acronym := ""
	for _, word := range words {
		acronym = acronym + strings.ToUpper(string([]rune(word)[0]))
	}
	return acronym
}

// BetterAbbreviate should have a comment documenting it.
func BetterAbbreviate(s string) (out string) {

	s = strings.Replace(s, "-", " ", -1)

	words := strings.Fields(s)

	for i := range words {
		out += string(words[i][0])
	}
	return strings.ToUpper(out)
}