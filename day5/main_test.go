package main

import "testing"

func TestPartOneFreshness(t *testing.T) {
	ranges := []RangePair{
		{3, 5, false},
		{10, 14, false},
		{16, 20, false},
		{12, 18, false}}
	ingredients := []int64{1, 5, 8, 11, 17, 32}
	expected := 3

	got := checkPartOneFreshness(ranges, ingredients)
	if got != expected {
		t.Errorf("checkPartOneFreshness() = %d; want %d", got, expected)
	}
}

func TestFindAllIDsCount(t *testing.T) {

	ranges := []RangePair{
		{3, 5, false},
		{10, 14, false},
		{16, 20, false},
		{12, 18, false}}
	expected := int64(14)

	got := checkAllIDsCount(ranges)
	if got != expected {
		t.Errorf("checkAllIDsCount() = %d; want %d", got, expected)
	}
}
