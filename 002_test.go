package main

import "testing"

type testpair002 struct {
	limit    int
	expected int
}

var tests002 = []testpair002{
	{100, 44},
	{4000000, 4613732},
}

func TestProblem002iterative(t *testing.T) {
	for _, test := range tests002 {
		sum := problem002iterative(test.limit)
		if sum != test.expected {
			t.Errorf("problem002iterative(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func TestProblem002closure(t *testing.T) {
	for _, test := range tests002 {
		sum := problem002closure(test.limit)
		if sum != test.expected {
			t.Errorf("problem002closure(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func BenchmarkProblem002iterative(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		problem002iterative(in)
	}
}

func BenchmarkProblem002closure(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		problem002closure(in)
	}
}
