package main

import "testing"

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{10, 23},
	{1000, 233168},
}

func TestProblem1unefficient(t *testing.T) {
	for _, test := range tests {
		sum := problem1unefficient(test.limit)
		if sum != test.expected {
			t.Errorf("problem1unefficient(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func TestProblem1efficient(t *testing.T) {
	for _, test := range tests {
		sum := problem1efficient(test.limit)
		if sum != test.expected {
			t.Errorf("problem1unefficient(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func BenchmarkProblem1unefficient(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		problem1unefficient(in)
	}
}

func BenchmarkProblem1efficient(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		problem1efficient(in)
	}
}
