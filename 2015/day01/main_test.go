package main

import "testing"

func TestCalculateFloor(t *testing.T) {
	input := readFile("input.txt")

	t.Run("Floor and basement are correctly calculated", func(t *testing.T) {
		floor, basement := calculateFloor(input)

		correctFloor, correctBasement := 280, 1797

		if floor != correctFloor {
			t.Errorf("got %d want %d", correctFloor, floor)
		}

		if basement != correctBasement {
			t.Errorf("got %d want %d", basement, correctBasement)
		}
	})
}
