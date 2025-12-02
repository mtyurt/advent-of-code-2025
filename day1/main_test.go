package main

import "testing"

func TestProcessInstructions(t *testing.T) {
	tt := []struct {
		instructions []string
		expected     int
	}{
		{[]string{"L50", "L414"}, 5},
		{[]string{"R50", "R60"}, 1},
		{[]string{"L100", "R200", "L300"}, 6},
		{[]string{"R50", "R1000", "L12"}, 11},
		{[]string{"L31", "L19"}, 1},
		{[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 6},
		{[]string{"R26", "L489", "L87", "R80"}, 6},
		{[]string{"L51"}, 1},         // Cross 0 once, end at 99
		{[]string{"L50"}, 1},         // End at 0, no crossing during
		{[]string{"L150"}, 2},        // Cross 0 once at step 50, end at 0 at step 150
		{[]string{"L49", "L101"}, 2}, // Go to position 1, then L101 crosses 0 once and ends at 0
	}

	for _, tc := range tt {
		answer, err := processInstructions(tc.instructions)
		if err != nil {
			t.Fatalf("unexpected error for instructions %v: %v", tc.instructions, err)
		}
		if answer != tc.expected {
			t.Fatalf("for instructions %v, expected answer %d, got %d", tc.instructions, tc.expected, answer)
		}
	}
}
