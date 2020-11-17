// Package hamming calculates the hamming distance between two genetic sequences
package hamming

import (
	"errors"
)

// Distance returns the hamming distance for two genetic sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("inputs must have the same length")
	} else {
		counter := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter++
			}
		}
		return counter, nil
	}

}
