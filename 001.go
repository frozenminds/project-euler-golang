// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
// https://projecteuler.net/problem=1

package main

// Find the sum of all the multiples of 3 or 5 below given limit
func problem1unefficient(limit int) int {
	sum := 0

	for i := 0; i < limit; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	return sum
}

// Find the sum of all the multiples of 3 or 5 below given limit
// Efficient solution https://projecteuler.net/overview=001
func problem1efficient(limit int) int {
	sdb := func(n int) int {
		p := (limit - 1) / n
		return n * (p * (p + 1)) / 2
	}

	return sdb(3) + sdb(5) - sdb(15)
}
