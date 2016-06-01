package problem006

import "testing"

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{10, 2640},
	{100, 25164150},
}

func TestSumSquareDiff(t *testing.T) {
	for _, test := range tests {
		diff := SumSquareDiff(test.limit)
		if diff != test.expected {
			t.Errorf("SumSquareDiff(%d) = %d, WANT: %d", test.limit, diff, test.expected)
		}
	}
}
