// Package diffsquares deals with math related problems
package diffsquares

// SquareOfSum returns the square of the sum of values from 1 to x
func SquareOfSum(x int) int {
	sum := x * (x + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of all squares from 1 to x
func SumOfSquares(x int) int {
	return x * (x + 1) * (2 * x + 1) / 6
}

// Difference returns the difference between SquareOfSum and SumOfSquares
func Difference(x, y int) int {
	return SquareOfSum(x) - SumOfSquares(y)
}