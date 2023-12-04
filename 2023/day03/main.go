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
type Point Interval
type Engine [][]rune

// var sample = `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
// `

func main() {
	engine := parseInput("input.txt")

	// engine := make([][]rune, 0)
	// split := strings.Split(sample, "\n")

	// for _, s := range split {
	// 	engine = append(engine, []rune(s))
	// }

	p1 := part1(engine)
	fmt.Println("part1:", p1)

	p2 := part2(engine)
	fmt.Println("part2:", p2)
}

func part1(engine Engine) int {
	parts := make(map[Point]int)

	for i := range engine {
		for j := range engine[i] {
			c := engine[i][j]

			// is not a symbol
			if !isSymbol(c) {
				continue
			}

			fmt.Println("\n===", i+1, j+1)

			// check row above, current, and below for adjacent part numbers
			for ii := utils.Max(0, i-1); ii <= utils.Min(len(engine)-1, i+1); ii++ {
				intervals := findIntervals(engine[ii])
				for _, interval := range intervals {
					l := utils.Max(0, j-1)
					r := utils.Min(len(engine[0])-1, j+1)

					if isOverlapping(interval, Interval{l, r + 1}) {
						p := Point{ii, interval.start}
						n, _ := strconv.Atoi(string(engine[ii][interval.start:interval.end]))
						parts[p] = n
						fmt.Printf("(%d,%d) %d\n", ii, j, n)
					}
				}
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
			c := engine[i][j]

			// is not a symbol
			if c != '*' {
				continue
			}

			fmt.Println("\n===", i+1, j+1)

			// check row above, current, and below for adjacent part numbers
			numAdjacent := make(map[Point]int, 0)

			for ii := utils.Max(0, i-1); ii <= utils.Min(len(engine)-1, i+1); ii++ {
				intervals := findIntervals(engine[ii])
				for _, interval := range intervals {
					l := utils.Max(0, j-1)
					r := utils.Min(len(engine[0])-1, j+1)

					if isOverlapping(interval, Interval{l, r + 1}) {
						p := Point{ii, interval.start}
						n, _ := strconv.Atoi(string(engine[ii][interval.start:interval.end]))
						numAdjacent[p] = n

						fmt.Printf("(%d,%d) %d\n", ii, j, n)
					}
				}
			}

			if len(numAdjacent) == 2 {
				val := 1
				for _, v := range numAdjacent {
					val *= v
				}

				total += val
			}
		}
	}

	return total
}

func findIntervals(runes []rune) []Interval {
	intervals := make([]Interval, 0)

	start := -1
	end := -1

	for i := range runes {
		c := runes[i]

		if start < 0 && unicode.IsNumber(c) {
			start = i
			continue
		}

		if start >= 0 && !unicode.IsNumber(c) {
			end = i
			intervals = append(intervals, Interval{start, end})

			start = -1
			end = -1
		}

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
