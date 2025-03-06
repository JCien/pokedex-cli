package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	//Create a slice of test case structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "MickeY MoUse",
			expected: []string{"mickey", "mouse"},
		},
		{
			input:    "Chuck E Cheese",
			expected: []string{"chuck", "e", "cheese"},
		},
	}

	// Loop over the cases
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test.
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match: %v vs %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test.
			if word != expectedWord {
				t.Errorf("Words do not match: %v vs %v", word, expectedWord)
			}
		}
	}
}
