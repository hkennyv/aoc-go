package main

import (
	"strings"
	"testing"
)

func TestFindAjacents(t *testing.T) {
	s := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	engine := [][]rune{}

	for _, ss := range strings.Split(s, "\n") {
		engine = append(engine, []rune(ss))
	}

	tests := []struct {
		x, y     int
		expected map[Point]int
	}{
		{
			1, 3, map[Point]int{
				{0, 0}: 467,
				{2, 2}: 35,
			},
		},
		{
			3, 6, map[Point]int{
				{2, 6}: 633,
			},
		},
		{
			4, 3, map[Point]int{
				{4, 0}: 617,
			},
		},
	}

	for _, test := range tests {
		actual := findAdjacents(engine, test.x, test.y)

		if len(actual) != len(test.expected) {
			t.Errorf("got %v, expected %v\n", actual, test.expected)
		}

		for k := range test.expected {
			if v, ok := actual[k]; ok {
				if v != test.expected[k] {
					t.Errorf("got %v, expected %v\n", v, test.expected[k])
				}
			} else {
				t.Errorf("expected %v, got %v\n", actual, test.expected)
			}
		}
	}
}

func TestFindIntervals(t *testing.T) {
	tests := []string{
		"...........820...358$....&..298........$....62...@....983..+..49.....254............784...*......................947.....................776",
		".....475.....................412.634*480..155....$824...845........409..367..........@..........*..498...............................893....",
	}

	expected := [][]Interval{
		{Interval{11, 14}, Interval{17, 20}, Interval{28, 31}, Interval{44, 46}, Interval{54, 57}, Interval{62, 64}, Interval{69, 72}, Interval{84, 87}, Interval{113, 116}, Interval{137, 140}},
		{Interval{5, 8}, Interval{29, 32}, Interval{33, 36}, Interval{37, 40}, Interval{42, 45}, Interval{50, 53}, Interval{56, 59}, Interval{67, 70}, Interval{72, 75}, Interval{99, 102}, Interval{133, 136}},
	}

	for i, test := range tests {
		intervals := findIntervals([]rune(test))

		for j, actual := range intervals {
			expected := expected[i][j]

			if actual.start != expected.start || actual.end != expected.end {
				t.Errorf("got %v, expected %v\n", actual, expected)
			}
		}
	}
}

func TestIsOverlapping(t *testing.T) {

	tests := []struct {
		i1, i2   Interval
		expected bool
	}{
		{Interval{1, 3}, Interval{2, 4}, true},
		{Interval{2, 5}, Interval{1, 3}, true},
		{Interval{5, 6}, Interval{1, 5}, false},
		{Interval{10, 14}, Interval{1, 4}, false},
	}

	for _, test := range tests {
		actual := isOverlapping(test.i1, test.i2)

		if actual != test.expected {
			t.Errorf("%v and %v got %v, expected %v\n", test.i1, test.i2, actual, test.expected)
		}
	}
}
