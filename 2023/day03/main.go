package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/hkennyv/aoc-go/utils"
)

type Interval struct {
	start int
	end   int
}
type Point struct {
	row int
	col int
}
type Engine [][]rune

func main() {
	engine := parseInput("input.txt")

	p1 := part1(engine)
	fmt.Println("part1:", p1)

	p2 := part2(engine)
	fmt.Println("part2:", p2)
}

func part1(engine Engine) int {
	parts := make(map[Point]int)

	for i := range engine {
		for j := range engine[i] {
			ch := engine[i][j]

			if !isSymbol(ch) {
				continue
			}

			adjacents := findAdjacents(engine, i, j)
			for k, v := range adjacents {
				parts[k] += v
			}
		}
	}

	total := 0

	for _, n := range parts {
		total += n
	}

	return total
}

func part2(engine Engine) int {
	total := 0

	for i := range engine {
		for j := range engine[i] {
			ch := engine[i][j]

			// is not a gear
			if ch != '*' {
				continue
			}

			adjacents := findAdjacents(engine, i, j)

			if len(adjacents) == 2 {
				val := 1
				for _, v := range adjacents {
					val *= v
				}

				total += val
			}
		}
	}

	return total
}

func findAdjacents(engine Engine, row, col int) map[Point]int {
	adjacents := make(map[Point]int, 0)

	// look in row above, current, and below the symbol
	//
	// find the intervals of all the numbers and see if it overlaps with the
	// column interval of the symbol

	// vertical bounds
	vstart := utils.Max(0, row-1)
	vend := utils.Min(len(engine)-1, row+1)

	for i := vstart; i <= vend; i++ {
		intervals := findIntervals(engine[i])

		for _, numInterval := range intervals {

			// horizontal bounds
			hstart := utils.Max(0, col-1)
			hend := utils.Min(len(engine[0])-1, col+1)

			// interval is left-inclusive, right non-inclusive
			symbolInteval := Interval{hstart, hend + 1}

			if isOverlapping(numInterval, symbolInteval) {
				p := Point{i, numInterval.start}
				n, _ := strconv.Atoi(string(engine[i][numInterval.start:numInterval.end]))

				adjacents[p] = n
			}
		}
	}

	return adjacents
}

func findIntervals(runes []rune) []Interval {
	intervals := make([]Interval, 0)

	start := -1
	end := -1

	for i := range runes {
		ch := runes[i]

		// start
		if start < 0 && unicode.IsNumber(ch) {
			start = i
			continue
		}

		// end of num in middle of line
		if start >= 0 && !unicode.IsNumber(ch) {
			end = i
			intervals = append(intervals, Interval{start, end})

			start = -1
			end = -1
		}

		// end of num end of the line
		if start >= 0 && i == len(runes)-1 {
			end = i + 1
			intervals = append(intervals, Interval{start, end})
		}
	}

	return intervals
}

func isSymbol(r rune) bool {
	return !unicode.IsNumber(r) && r != '.'
}

func isOverlapping(i1, i2 Interval) bool {
	return (i1.start >= i2.start && i1.start < i2.end) || (i2.start >= i1.start && i2.start < i1.end)
}

func parseInput(fn string) Engine {
	b, _ := os.ReadFile(fn)
	split := strings.Split(strings.TrimSpace(string(b)), "\n")

	runes := make([][]rune, 0)

	for _, s := range split {
		runes = append(runes, []rune(s))
	}

	return runes
}
