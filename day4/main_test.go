package main

import "testing"

func TestProcessInput(t *testing.T) {

	input := `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	expected := 13
	if got := processInput(input); got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
