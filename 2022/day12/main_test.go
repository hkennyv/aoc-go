package main

import "testing"

var sample [][]rune = [][]rune{
	[]rune("Sabqponm"), []rune("abcryxxl"), []rune("accszExk"), []rune("acctuvwj"), []rune("abdefghi"),
}

func TestFindPath(t *testing.T) {
	expected := 31
	start, err := findPoint('S', sample)
	if err != nil {
		t.Error(err)
	}
	end, err := findPoint('E', sample)
	if err != nil {
		t.Error(err)
	}

	sample[start.X][start.Y] = 'a'
	sample[end.X][end.Y] = 'z'

	actual, err := findPath(start, end, sample)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected %d, got %d\n", expected, actual)
	}
}
