package main

import "testing"

func TestLowestNumberHash(t *testing.T) {

	type Case struct {
		input    string
		expected int
	}

	var cases = []Case{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	t.Run("finds the lowest number that produces hash with five leadings zeros", func(t *testing.T) {
		for _, c := range cases {
			got := lowestNumberHash(c.input, "00000")

			if got != c.expected {
				t.Errorf("got %d want %d", got, c.expected)
			}
		}
	})
}
