// Package raindrops returns plingplangplong depending on factors of input
package raindrops

import (
	"strconv"
)

// Convert takes in an int input and returns a string based on factors
func Convert(input int) string {
	output := ""
	if input % 3 == 0 {
		output += "Pling"
	}
	if input % 5 == 0 {
		output += "Plang"
	}
	if input % 7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(input)
	}
	return output
}
