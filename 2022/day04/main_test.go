package main

import (
	"strings"
	"testing"
)

const sample = `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

var expected = [][]int{
	{2, 4, 6, 8},
	{2, 3, 4, 5},
	{5, 7, 7, 9},
	{2, 8, 3, 7},
	{6, 6, 4, 6},
	{2, 6, 4, 8},
}

func TestParseLine(t *testing.T) {
	input := strings.Split(strings.TrimSpace(sample), "\n")
	res := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = parseLine(input[i])
	}

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			if res[i][j] != expected[i][j] {
				t.Errorf("%d should be %d\n", res[i][j], expected[i][j])
			}
		}
	}
}

func TestFullyOverlaps(t *testing.T) {
	answers := []bool{false, false, false, true, true, false}

	for i, ans := range answers {
		actual := fullyOverlaps(expected[i])
		if actual != ans {
			t.Errorf("got %v for %v, should be %v\n", actual, expected[i], ans)
		}
	}
}

func TestPartialOverlaps(t *testing.T) {
	answers := []bool{false, false, true, true, true, true}

	for i, ans := range answers {
		actual := partialOverlaps(expected[i])
		if actual != ans {
			t.Errorf("got %v for %v, should be %v\n", actual, expected[i], ans)
		}
	}
}
