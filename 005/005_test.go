package problem005

import "testing"

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{10, 2520},
	{20, 232792560},
}

func TestSmallestMultiple(t *testing.T) {
	for _, test := range tests {
		multiple := SmallestMultiple(test.limit)
		if multiple != test.expected {
			t.Errorf("SmallestMultiple(%d) = %d, WANT: %d", test.limit, multiple, test.expected)
		}
	}
}
