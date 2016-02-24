package problem002

import "testing"

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{100, 44},
	{4000000, 4613732},
}

func TestIterative(t *testing.T) {
	for _, test := range tests {
		sum := Iterative(test.limit)
		if sum != test.expected {
			t.Errorf("Iterative(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func TestClosure(t *testing.T) {
	for _, test := range tests {
		sum := Closure(test.limit)
		if sum != test.expected {
			t.Errorf("Closure(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func BenchmarkIterative(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		Iterative(in)
	}
}

func BenchmarkClosure(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		Closure(in)
	}
}
