package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Println("hello world")

	f, err := os.ReadFile("input.txt")
	check(err)

	body := strings.TrimSpace(string(f))
	total := 0
	for line := range strings.SplitSeq(body, "\n") {
		if line == "" {
			continue
		}
		total += maxNDigitJoltage(line, 12)
	}
	fmt.Println("total:", total)
}

func maxJoltage(in string) int {

	firstLargest, secondLargest := 0, 0
	largestIndex := 0
	for i := 0; i < len(in)-1; i++ {
		digit, err := strconv.Atoi(string(in[i]))
		check(err)

		if digit > firstLargest {
			firstLargest = digit
			largestIndex = i
		}
	}
	for i := largestIndex + 1; i < len(in); i++ {
		digit, err := strconv.Atoi(string(in[i]))
		check(err)

		if digit > secondLargest {
			secondLargest = digit
		}
	}

	return firstLargest*10 + secondLargest

}

func getDigit(in string, i int) int {

	digit, err := strconv.Atoi(string(in[i]))
	check(err)
	return digit

}

func maxNDigitJoltage(in string, lookingFor int) int {
	totalDigits := len(in)
	// fmt.Println("checking", in, "totalDigits", totalDigits)

	digits := make([]int, lookingFor)
	beginIndex := 0
	total := 0
	for i := range lookingFor {
		firstLargestInWindow := 0
		largestIndexInWindow := 0
		// fmt.Printf("searching from: %d until: %d for %dth index\n", beginIndex, totalDigits-12+i, i)
		for j := beginIndex; j <= totalDigits-lookingFor+i; j++ {
			digit := getDigit(in, j)
			if digit > firstLargestInWindow {
				firstLargestInWindow = digit
				largestIndexInWindow = j
			}
		}
		// fmt.Printf("digit %d: %d, found at: %d\n", i, firstLargestInWindow, largestIndexInWindow)
		digits[i] = firstLargestInWindow
		beginIndex = largestIndexInWindow + 1
		total = total*10 + digits[i]
	}

	return total
}
