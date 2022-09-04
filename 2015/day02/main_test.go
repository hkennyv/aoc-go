package main

import "testing"

func TestMakePresent(t *testing.T) {
	s := "2x3x4"

	t.Run("test present is correct", func(t *testing.T) {
		p := MakePresent(s)

		expectSides := []int{2, 3, 4}

		for i, side := range p.sides {
			if side != expectSides[i] {
				t.Errorf("got %d want %d", side, expectSides[i])
			}
		}
	})

	t.Run("calculates square feet correctly", func(t *testing.T) {
		p := MakePresent(s)
		got := p.CalculateSquareFeet()
		want := 58

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("calculates ribbon correctly", func(t *testing.T) {
		p := MakePresent(s)
		got := p.CalcualteRibbon()
		want := 34

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
