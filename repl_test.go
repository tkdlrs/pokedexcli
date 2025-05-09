package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// ...
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander BuLbAsAuR PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	//
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Word does not match expected word: Received: %v, Expected %v", actual, c.expected)
			}
		}
	}
}
