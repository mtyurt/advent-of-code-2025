package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangePair struct {
	begin int64
	end   int64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Println("vim-go")

	f, _ := os.ReadFile("input.txt")

	body := strings.TrimSpace(string(f))
	lines := strings.Split(body, "\n")

	ranges := make([]RangePair, 0)
	lastIndex := 0
	for i, l := range lines {

		if strings.TrimSpace(l) == "" {
			lastIndex = i
			break
		}
		splitted := strings.Split(l, "-")
		begin, err := strconv.ParseInt(splitted[0], 10, 64)
		check(err)
		end, err := strconv.ParseInt(splitted[1], 10, 64)
		check(err)

		ranges = append(ranges, RangePair{begin, end})
	}

	ingredients := make([]int64, 0)
	for _, l := range lines[lastIndex+1:] {
		id, err := strconv.ParseInt(l, 10, 64)
		check(err)
		ingredients = append(ingredients, id)
	}

	total := checkPartOneFreshness(ranges, ingredients)
	fmt.Println(total)

}

func checkPartOneFreshness(ranges []RangePair, ingredients []int64) int {
	total := 0
	for _, i := range ingredients {
	rangecheck:
		for _, j := range ranges {
			if i >= j.begin && i <= j.end {
				total++
				break rangecheck
			}
		}
	}
	return total
}
