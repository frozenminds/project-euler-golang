package problem003

import "testing"

type testpair struct {
	value    int64
	expected int64
}

var tests = []testpair{
	{600851475143, 6857},
}

func TestPrimeMax(t *testing.T) {
	for _, test := range tests {
		prime := PrimeMax(test.value)
		if prime != test.expected {
			t.Errorf("PrimeMax(%d) = %d, WANT: %d", test.value, prime, test.expected)
		}
	}
}
