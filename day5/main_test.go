package main

import "testing"

func TestPartOneFreshness(t *testing.T) {
	ranges := []RangePair{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18}}
	ingredients := []int64{1, 5, 8, 11, 17, 32}
	expected := 3

	got := checkPartOneFreshness(ranges, ingredients)
	if got != expected {
		t.Errorf("checkPartOneFreshness() = %d; want %d", got, expected)
	}
}
