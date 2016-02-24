package problem001

import "testing"

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{10, 23},
	{1000, 233168},
}

func TestUnefficient(t *testing.T) {
	for _, test := range tests {
		sum := Unefficient(test.limit)
		if sum != test.expected {
			t.Errorf("Unefficient(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func TestEfficient(t *testing.T) {
	for _, test := range tests {
		sum := Efficient(test.limit)
		if sum != test.expected {
			t.Errorf("Efficient(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func BenchmarkUnefficient(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		Unefficient(in)
	}
}

func BenchmarkEfficient(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		Efficient(in)
	}
}
