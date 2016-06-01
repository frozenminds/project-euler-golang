// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
// https://projecteuler.net/problem=7

package problem007

// is prime number
func prime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// NthPrimeNumber
func NthPrimeNumber(number int) int {
	i := 1

	for number > 0 {
		i++

		if prime(i) {
			number--
		}
	}

	return i
}
