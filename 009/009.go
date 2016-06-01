// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
// https://projecteuler.net/problem=9

package problem009

// PythagoreanTriplet
func PythagoreanTriplet(sum int) int {

	for a := 1; a <= sum; a++ {
		for b := (a + 1); b <= sum; b++ {
			c := sum - a - b

			if a*a+b*b == c*c && a+b+c == sum {
				return a * b * c
			}
		}

	}
	return 0
}
