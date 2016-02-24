// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// https://projecteuler.net/problem=3

package problem003

import "fmt"

// Check if number is prime
func prime(n int) bool {
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Find maximum prime factor
func PrimeMax(n int) int {

	primes := []int{}

	f := func(primes []int, n int) bool {

		res := 1
		for _, v := range primes {
			res *= v
			if res > n {
				return false
			}
			if res == n {
				return true
			}
		}

		return false
	}

	for i := 2; i <= n; i++ {
		if prime(i) {
			primes = append(primes, i)
			if f(primes, n) {
				fmt.Println("found max: ", i)
				return i
			}
		}
	}

	return 2
}
