// The sum of the squares of the first ten natural numbers is, 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is, (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
// https://projecteuler.net/problem=6

package problem006

// 1^2 + 2^2 + ... + limit^2
func sumOfSquares(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum = sum + (i * i)
	}
	return sum
}

// (1 + 2 + ... + limit)^2
func squareOfSums(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum = sum + i
	}
	return sum * sum
}

// SumSquareDiff
func SumSquareDiff(limit int) int {
	return squareOfSums(limit) - sumOfSquares(limit)
}
