// Package luhn helps determines whether a number is valid based on the Luhn formula
package luhn

import (
	"strings"
	"strconv"
)

// Valid helps check if a number is a valid credit card number
func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	sum := 0
	size := len(num)
	if len(num) <= 1 {
		return false
	}
	if _, err := strconv.Atoi(num); err != nil {
		return false
	}
	
	for i := 0; i < size; i++ {
		pos := size - 1 - i
		digit, _ := strconv.Atoi(string(num[pos]))
		if i % 2 == 0 {
			sum += digit
		} else {
			if digit >= 5 {
				sum += digit * 2 - 9
			} else {
				sum += digit * 2
			}
		}
	}

	if sum % 10 == 0 {
		return true
	}
	return false
} 
