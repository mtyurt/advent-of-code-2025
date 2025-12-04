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
		total += maxJoltage(line)
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
