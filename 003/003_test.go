package problem003

import "testing"

type testpair struct {
	value    int
	expected int
}

var tests = []testpair{
	{13195, 29},
}

func TestPrimeMax(t *testing.T) {
	for _, test := range tests {
		prime := PrimeMax(test.value)
		if prime != test.expected {
			t.Errorf("PrimeMax(%d) = %d, WANT: %d", test.value, prime, test.expected)
		}
	}
}
