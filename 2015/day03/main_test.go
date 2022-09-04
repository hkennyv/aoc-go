package main

import "testing"

type Case struct {
	input     string
	expected1 int
	expected2 int
}

func TestDirections(t *testing.T) {

	var cases = []Case{
		{">", 2, 2},
		{"^>v<", 4, 3},
		{"^v^v^v^v^v", 2, 11},
	}

	t.Run("directions are calculated correctly for santa", func(t *testing.T) {
		for _, c := range cases {
			got := calculateHouses(c.input)
			want := c.expected1

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})

	t.Run("directions are calculated correctly for santa & robosanta", func(t *testing.T) {
		for _, c := range cases {
			got := calculateHousesRobo(c.input)
			want := c.expected2

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})
}
