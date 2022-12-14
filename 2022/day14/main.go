package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

type Coordinate struct {
	col, row int
}

type Trace = []Coordinate

type Grid = [][]rune

func main() {
	traces := parseInput("input.txt")
	grid := makeGrid(traces, false)
	_, _, b := getBounds(grid)

	untilVoid := func(row int) bool {
		return row <= b
	}
	fmt.Println("part1:", pourSand(grid, untilVoid))

	grid = makeGrid(traces, true)
	untilSourceBlocked := func(row int) bool {
		return true
	}
	pourSand(grid, untilSourceBlocked)
}

func pourSand(g Grid, cond func(int) bool) int {
	count := 0
	col, row := 500, 0

	// when row goes under b, we have gone into the void...
	for cond(row) {

		if g[row][col] != 0 {

			// check diagonal left & right
			if g[row][col-1] != 0 && g[row][col+1] != 0 {
				if row-1 == -1 {
					printGrid(g)
					fmt.Println("OOB!", count)
				}
				g[row-1][col] = 'o'
				count++
				col = 500
				row = 0
				// move left
			} else if g[row][col-1] == 0 {
				col--
				// move right
			} else if g[row][col+1] == 0 {
				col++
			}

			// keep falling
		} else {
			row++
		}
	}

	return count
}

func makeGrid(ts []Trace, withFloor bool) Grid {
	// make 1000x1000 grid
	N := 1000
	grid := make(Grid, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]rune, N)
	}

	// fill in grid
	for _, t := range ts {
		for i := 1; i < len(t); i++ {
			c1, c2 := t[i-1], t[i]
			startc, endc := utils.Min(c1.col, c2.col), utils.Max(c1.col, c2.col)
			for j := startc; j <= endc; j++ {
				startr, endr := utils.Min(c1.row, c2.row), utils.Max(c1.row, c2.row)
				for k := startr; k <= endr; k++ {
					grid[k][j] = '#'
				}
			}
		}
	}

	if withFloor {
		_, _, b := getBounds(grid)
		for i := 0; i < N; i++ {
			grid[b+2][i] = '#'
		}
	}

	return grid
}

func getBounds(g Grid) (int, int, int) {
	l, r, b := math.MaxInt, 0, 0

	// find mins and maxes
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			curr := g[i][j]

			if curr == '#' {
				l = utils.Min(l, j)
				r = utils.Max(r, j)
				b = utils.Max(b, i)
			}
		}
	}

	return l, r, b
}

func printGrid(g Grid) {
	l, r, b := getBounds(g)

	for i := l; i <= r; i++ {
		if i == 500 {
			fmt.Printf("x")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("\n")
	for i := 0; i <= b; i++ {
		for j := l; j <= r; j++ {
			if g[i][j] != 0 {
				fmt.Printf("%c", g[i][j])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func parseInput(fp string) []Trace {
	b, _ := os.ReadFile(fp)
	split := strings.Split(strings.TrimSpace(string(b)), "\n")
	ret := make([]Trace, len(split))
	for i := 0; i < len(split); i++ {
		ss := strings.Split(split[i], " -> ")
		ret[i] = make(Trace, len(ss))
		for j := 0; j < len(ss); j++ {
			cs := strings.Split(ss[j], ",")
			col, _ := strconv.Atoi(cs[0])
			row, _ := strconv.Atoi(cs[1])
			ret[i][j] = Coordinate{col, row}
		}
	}
	return ret
}
