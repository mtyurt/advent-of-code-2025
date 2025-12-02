package main

import (
	"slices"
	"testing"
)

func TestInvalidIDCheck(t *testing.T) {
	ttest := []struct {
		id string
		ok bool
	}{
		{"", false},
		{"11", true},
		{"22", true},
		{"111", true},
		{"35", false},
		{"999", true},
		{"1010", true},
		{"446", false},
		{"010", false},
		{"565656", true},
		{"824824824", true},
		{"123456123456", true},
		{"2121212121", true},
		{"38593859", true},
		{"38593857", false},
		{"3859385", false},
		{"1188511885", true},
		{"446446", true},
	}
	for _, tc := range ttest {
		got := isIDInvalid(tc.id)
		if got != tc.ok {
			t.Errorf("isIDInvalid(%q) = %v; want %v", tc.id, got, tc.ok)
		}
	}

}

func TestCheckIDRange(t *testing.T) {
	ttest := []struct {
		begin int64
		end   int64
		ids   []int64
	}{
		{11, 22, []int64{11, 22}},
		{95, 115, []int64{99, 111}},
		{998, 1012, []int64{999, 1010}},
		{1188511880, 1188511890, []int64{1188511885}},
		{1698522, 1698528, []int64{}},
		{222220, 222224, []int64{222222}},
		{446443, 446449, []int64{446446}},
		{565653, 565659, []int64{565656}},
		{824824821, 824824827, []int64{824824824}},
		{2121212118, 2121212124, []int64{2121212121}},
		{38593856, 38593862, []int64{38593859}},
	}
	for _, tc := range ttest {
		got := checkIDRange(tc.begin, tc.end)
		if !slices.Equal(got, tc.ids) {
			t.Errorf("checkIDRange(%d, %d) = %v; want %v", tc.begin, tc.end, got, tc.ids)
		}
	}

}
