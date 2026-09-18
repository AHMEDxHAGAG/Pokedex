package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "  HELLO  WORLD  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "",
			expected: []string{},
		}, {
			input:    "               H                ",
			expected: []string{"h"},
		},
	}
	for no, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("test %d: actual Len: %d doesnt match the expected %d", no, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("test %d: actual word: %s doesnt match the expected word %s", no, actual, c.expected)
			}
		}
	}
}
