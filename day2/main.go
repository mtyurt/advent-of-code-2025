package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

func isIDInvalid(id string) bool {
	fmt.Println("checking id:", id)
	len := len(id)
	if len == 0 || len%2 != 0 {
		return false
	}

	for i := range len / 2 {
		fmt.Printf("checking %s %s\n", string(id[i]), string(id[len-i-1]))
		if id[i] != id[i+len/2] {
			return false
		}
	}

	return true
}
