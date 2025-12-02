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
	ranges := strings.Split(body, ",")
	total := int64(0)
	for _, r := range ranges {
		s := strings.Split(r, "-")
		begin, err := strconv.ParseInt(s[0], 10, 64)
		check(err)
		end, err := strconv.ParseInt(s[1], 10, 64)
		check(err)
		invalids := checkIDRange(begin, end)
		for _, id := range invalids {
			total += id
		}
	}

	fmt.Println(total)

}

func isIDInvalid(id string) bool {
	len := len(id)
	if len == 0 || len%2 != 0 {
		return false
	}

	for i := range len / 2 {
		if id[i] != id[i+len/2] {
			return false
		}
	}

	return true
}

func checkIDRange(begin, end int64) []int64 {
	ids := []int64{}
	for id := begin; id <= end; id++ {
		if isIDInvalid(fmt.Sprintf("%d", id)) {
			ids = append(ids, id)
		}
	}
	return ids
}
