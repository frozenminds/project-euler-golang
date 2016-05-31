// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// https://projecteuler.net/problem=4

package problem004

import (
	"bytes"
	"log"
	"strconv"
)

// get largest number by digits count
func getLargestNumber(digits int) int {

	var buffer bytes.Buffer
	for i := 0; i < digits; i++ {
		buffer.WriteString("9")
	}

	num, err := strconv.Atoi(buffer.String())
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// get smallest number by digits count
func getSmallestNumber(digits int) int {

	var buffer bytes.Buffer
	buffer.WriteString("1")
	for i := 1; i < digits; i++ {
		buffer.WriteString("0")
	}

	num, err := strconv.Atoi(buffer.String())
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// reverse a number
func reverseNumber(number int) int {
	reverse := 0

	for number > 0 {
		reverse = (reverse * 10) + (number % 10)
		number = number / 10
	}
	return reverse
}

func palindrome(number int) bool {
	if number == reverseNumber(number) {
		return true
	}
	return false
}

// PalindromeLargest
func PalindromeLargest(digits int) int {

	result := 0
	largest := getLargestNumber(digits)
	smallest := getSmallestNumber(digits)

	for a := largest; a > smallest; a-- {
		for b := largest; b > smallest; b-- {
			product := a * b

			if palindrome(product) && product > result {
				result = product
			}

		}
	}

	return result
}
