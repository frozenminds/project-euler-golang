package problem004

import "testing"

type testpair struct {
	digits   int
	expected int
}

var tests = []testpair{
	{2, 9009},
	{3, 906609},
}

func TestPalindromeLargest(t *testing.T) {
	for _, test := range tests {
		palindrome := PalindromeLargest(test.digits)
		if palindrome != test.expected {
			t.Errorf("PalindromeLargest(%d) = %d, WANT: %d", test.digits, palindrome, test.expected)
		}
	}
}
