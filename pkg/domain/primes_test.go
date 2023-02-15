package domain

import "testing"

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"negative numbers", -1, false},
		{"0", 0, false},
		{"2", 2, true},
		{"even number", 4, false},
		{"prime number", 11, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if IsPrime(tc.input) != tc.expected {
				t.Errorf("expected (%d) to be %v, but it should have been %v", tc.input, tc.expected, !tc.expected)
			}
		})
	}
}
