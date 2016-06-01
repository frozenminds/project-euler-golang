package problem009

import "testing"

type testpair struct {
	number   int
	expected int
}

var tests = []testpair{
	{1000, 31875000},
}

func TestPythagoreanTriplet(t *testing.T) {
	for _, test := range tests {
		product := PythagoreanTriplet(test.number)
		if product != test.expected {
			t.Errorf("PythagoreanTriplet(%d) = %d, WANT: %d", test.number, product, test.expected)
		}
	}
}
