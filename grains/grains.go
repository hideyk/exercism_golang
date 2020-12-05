// Package grains counts the number of grains on a particular grid
package grains

import (
	"math"
	"errors"
)

// Square returns the number of grains on a chess grid and whether its feasible
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("Stupid error")
	}
	
	return uint64(math.Pow(2, float64(num-1))), nil
}

//Total returns the total number of of grains on the chessboard
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		sum += uint64(math.Pow(2, float64(i-1)))
	}
	return sum
}
