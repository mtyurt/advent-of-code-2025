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
	// f, err := os.ReadFile("test.txt")
	check(err)

	fmt.Println(string(f))

	instructions := strings.Split(string(f), "\n")

	answer, err := processInstructions(instructions)
	check(err)
	fmt.Println("Answer:", answer)
}

func processInstructions(instructions []string) (int, error) {
	position := 50
	zeroCount := 0
	for i, instruction := range instructions {
		if instruction == "" {
			continue
		}
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return 0, fmt.Errorf("invalid distance in instruction %s at line %d", instruction, i)
		}
		if instruction[0] == 'L' {
			distance = -distance
		}
		position += distance + 100
		position = position % 100

		if position == 0 {
			zeroCount = zeroCount + 1
		}
		fmt.Println("Instruction:", instruction, " New Position:", position, " Distance:", distance)
	}

	return zeroCount, nil
}
