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
