package main

import "testing"

func TestInvalidIDCheck(t *testing.T) {
	ttest := []struct {
		id string
		ok bool
	}{
		{"", false},
		{"11", true},
		{"22", true},
		{"35", false},
		{"1010", true},
		{"446", false},
		{"010", false},
		{"123456123456", true},
		{"38593859", true},
		{"38593857", false},
		{"3859385", false},
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
		{95, 115, []int64{99}},
		{998, 1012, []int64{1010}},
		{1188511880, 1188511890, []int64{1188511885}},
		{222220, 222224, []int64{222222}},
		{446443, 446449, []int64{446446}},
		{38593856, 38593862, []int64{38593859}},
	}
	for _, tc := range ttest {
		got := checkIDRange(tc.begin, tc.end)
		if !slices.Equal(got, tc.ids) {
			t.Errorf("checkIDRange(%d, %d) = %v; want %v", tc.begin, tc.end, got, tc.ids)
		}
	}

}
