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
		result := maxNDigitJoltage(tc.input, 2)
		if result != tc.expected {
			t.Errorf("maxJoltage(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func TestMaxTwelveDigitJoltage(t *testing.T) {

	ttest := []struct {
		input    string
		expected int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tc := range ttest {
		result := maxNDigitJoltage(tc.input, 12)
		if result != tc.expected {
			t.Errorf("maxTwelveDigitJoltage(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
