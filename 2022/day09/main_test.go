package main

import (
	"testing"
)

var sample [][]string = [][]string{
	{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"},
	{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"},
}

func TestSimulate(t *testing.T) {
	expected := []int{13, 36}
	actuals := []int{simulate(sample[0], 2), simulate(sample[1], 10)}

	for i := 0; i < len(expected); i++ {
		if actuals[i] != expected[i] {
			t.Errorf("[%d] %d should be %d\n", i, actuals[i], expected[i])
		}
	}
}
