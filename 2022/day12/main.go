package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

type Point struct {
	X, Y int
}

func main() {
	grid := parseInput("input.txt")
	start, _ := findPoint('S', grid)
	end, _ := findPoint('E', grid)

	// get rid of S and E
	grid[start.X][start.Y] = 'a'
	grid[end.X][end.Y] = 'z'

	n, err := findPath(start, end, grid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("part1:", n)

	starts := make([]Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'a' {
				starts = append(starts, Point{i, j})
			}
		}
	}

	m := math.MaxInt
	for _, s := range starts {
		n, err := findPath(s, end, grid)
		if err != nil {
			continue
		}
		m = utils.Min(m, n)
	}
	fmt.Println("part2:", m)
}

func findPath(start, end Point, grid [][]rune) (int, error) {
	steps := 0
	visited := make(map[Point]struct{}, 0)
	visited[start] = struct{}{}
	q := []Point{start}

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			p := q[0]
			q = q[1:]

			if p == end {
				return steps, nil
			}

			neighbors := getSurrounding(p, grid)
			for _, n := range neighbors {
				if _, ok := visited[n]; ok {
					continue
				}
				visited[n] = struct{}{}
				q = append(q, n)
			}
		}
		steps++
	}

	return -1, fmt.Errorf("cannot find a path from %v to %v", start, end)
}

func getSurrounding(p Point, grid [][]rune) []Point {
	ret := make([]Point, 0)

	for _, d := range []Point{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	} {
		x := p.X + d.X
		y := p.Y + d.Y

		// in bounds and is one step up
		if x >= 0 &&
			x < len(grid) &&
			y >= 0 &&
			y < len(grid[0]) &&
			// grid[x][y]-grid[p.X][p.Y] >= -1 && 		// **at most one higher** not lower i guess
			grid[x][y]-grid[p.X][p.Y] <= 1 {
			ret = append(ret, Point{x, y})
		}
	}

	return ret
}

func findPoint(l rune, grid [][]rune) (Point, error) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == l {
				return Point{x, y}, nil
			}
		}
	}

	return Point{}, fmt.Errorf("cannot find %c in grid", l)
}

func parseInput(fp string) [][]rune {
	b, _ := os.ReadFile(fp)
	ss := strings.Split(strings.TrimSpace(string(b)), "\n")
	res := make([][]rune, len(ss))
	for i := 0; i < len(ss); i++ {
		res[i] = []rune(ss[i])
	}
	return res
}
