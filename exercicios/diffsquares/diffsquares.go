package diffsquares

import (
	"math"
)

// SquareOfSum calcula o quadrado da soma dos primeiros n números naturais
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares calcula a soma dos quadrados dos primeiros n números naturais
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calcula a diferença entre o quadrado da soma e a soma dos quadrados
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
