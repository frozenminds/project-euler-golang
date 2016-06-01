package problem007

import "testing"

type testpair struct {
	number   int
	expected int
}

var tests = []testpair{
	{6, 13},
	{10001, 104743},
}

func TestNthPrimeNumber(t *testing.T) {
	for _, test := range tests {
		prime := NthPrimeNumber(test.number)
		if prime != test.expected {
			t.Errorf("NthPrimeNumber(%d) = %d, WANT: %d", test.number, prime, test.expected)
		}
	}
}
