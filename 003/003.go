// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
// https://projecteuler.net/problem=3

package problem003

// get max
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// PrimeMax - Find maximum prime factor
func PrimeMax(n int64) int64 {

	factor := n
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			factor = max(i, n/i)
			break
		}
	}
	if factor == n {
		return n
	}
	return PrimeMax(factor)
}
