package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	f, _ := os.ReadFile("input.txt")

	body := strings.TrimSpace(string(f))
	fmt.Println(processInput(body))
}

func processInput(input string) int {

	matrix := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, len(line))
		for j, char := range line {
			row[j] = char
		}
		matrix = append(matrix, row)
	}
	total := 0
	for i := range len(matrix) {
		for j := 0; j < len(matrix[i]); j++ {
			char := matrix[i][j]
			if char != '@' {
				continue
			}
			occupiedNeighbors := 0

			// fmt.Printf("i: %d j: %d | ", i, j)
			if i > 0 && matrix[i-1][j] == '@' {
				occupiedNeighbors++
			}
			if i < len(matrix)-1 && matrix[i+1][j] == '@' {
				occupiedNeighbors++
			}
			if j < len(matrix[i])-1 && matrix[i][j+1] == '@' {
				occupiedNeighbors++
			}
			if j > 0 && matrix[i][j-1] == '@' {
				occupiedNeighbors++
			}

			if i > 0 && j > 0 && matrix[i-1][j-1] == '@' {
				occupiedNeighbors++
			}
			if i < len(matrix)-1 && j < len(matrix[i])-1 && matrix[i+1][j+1] == '@' {
				occupiedNeighbors++
			}
			if j < len(matrix[i])-1 && i > 0 && matrix[i-1][j+1] == '@' {
				occupiedNeighbors++
			}
			if j > 0 && i < len(matrix)-1 && matrix[i+1][j-1] == '@' {
				occupiedNeighbors++
			}

			if occupiedNeighbors < 4 {
				total++
			}

		}
		fmt.Printf("Line: %d total: %d\n", i, total)
	}
	return total
}
