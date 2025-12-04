package main

import "testing"

func TestMaxJoltage(t *testing.T) {

	ttest := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, tc := range ttest {
		result := maxJoltage(tc.input)
		if result != tc.expected {
			t.Errorf("maxJoltage(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
