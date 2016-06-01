// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// https://projecteuler.net/problem=5

package problem005

func divisible(number, limit int) bool {
	for i := 2; i <= limit; i++ {
		if number%i != 0 {
			return false
		}
	}

	return true
}

// SmallestMultiple
func SmallestMultiple(limit int) int {
	number := 2
	for {
		if divisible(number, limit) {
			return number
		}
		number = number + 2
	}
}
