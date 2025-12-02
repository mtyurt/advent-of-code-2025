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
	// fmt.Println("checking id:", id)
	len := len(id)
	if len == 0 {
		return false
	}
	for i := 0; i < len; i++ {
		begin, end := 0, i
		// fmt.Println("i:", i)
		// fmt.Println("begin:", begin, "end:", end)
		if end+1 >= len {
			return false
		}
		patternMatch := true
		for j := end + 1; j < len; {
			// fmt.Println("j:", j)
		innest:
			for k := begin; k <= end; k++ {
				// fmt.Println("k:", k)
				if j+k >= len {
					patternMatch = false
					break innest
				}

				// fmt.Println("id[k]:", id[k], "id[j+k]:", id[j+k])
				if id[k] != id[j+k] {
					patternMatch = false
					break innest
				}
			}
			j += (end - begin + 1)
		}
		if patternMatch {
			return true
		}
	}
	return false
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
