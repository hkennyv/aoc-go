package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

func main() {
	input := parseInput("input.txt")
	trees := findTrees(input)
	score := findScenicSpot(input)

	fmt.Println("part1:", trees)
	fmt.Println("part2:", score)
}

func findTrees(trees [][]int) int {
	seen := make(map[Coordinate]struct{}, 0)
	rows, cols := len(trees), len(trees[0])

	// add all border trees
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				c := Coordinate{i, j}
				seen[c] = struct{}{}
			}
		}
	}

	// check each row forward and backward pass
	for i := 0; i < rows; i++ {
		lheight := trees[i][0]
		rheight := trees[i][cols-1]

		for j := 0; j < cols; j++ {
			lcell := trees[i][j]
			rcell := trees[i][cols-1-j]

			if lcell > lheight {
				lheight = lcell
				c := Coordinate{i, j}
				seen[c] = struct{}{}
			}

			if rcell > rheight {
				rheight = rcell
				c := Coordinate{i, cols - 1 - j}
				seen[c] = struct{}{}
			}
		}
	}

	// check each col top down and bottom up pass
	for i := 0; i < cols; i++ {
		theight := trees[0][i]
		bheight := trees[rows-1][i]

		for j := 0; j < rows; j++ {
			tcell := trees[j][i]
			bcell := trees[rows-1-j][i]

			if tcell > theight {
				theight = tcell
				c := Coordinate{j, i}
				seen[c] = struct{}{}
			}

			if bcell > bheight {
				bheight = bcell
				c := Coordinate{rows - 1 - j, i}
				seen[c] = struct{}{}
			}
		}
	}

	return len(seen)
}

func findScenicSpot(trees [][]int) int {
	scores := make(map[Coordinate]int, 0)
	rows, cols := len(trees), len(trees[0])

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			c := Coordinate{i, j}
			scores[c] = calculateScenicScore(c, trees)
		}
	}

	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}

	return max
}

func calculateScenicScore(c Coordinate, trees [][]int) int {
	// check top view
	top := 0
	for i := c.x; i >= 0; i-- {
		if trees[c.x][c.y] <= trees[i][c.y] && i != c.x {
			top++
			break
		}

		if trees[c.x][c.y] > trees[i][c.y] {
			top++
		}
	}

	// check bottom view
	bottom := 0
	for i := c.x; i < len(trees); i++ {
		if trees[c.x][c.y] <= trees[i][c.y] && i != c.x {
			bottom++
			break
		}

		if trees[c.x][c.y] > trees[i][c.y] {
			bottom++
		}
	}

	// check left view
	left := 0
	for i := c.y; i >= 0; i-- {
		if trees[c.x][c.y] <= trees[c.x][i] && i != c.y {
			left++
			break
		}

		if trees[c.x][c.y] > trees[c.x][i] {
			left++
		}
	}

	// check right view
	right := 0
	for i := c.y; i < len(trees[0]); i++ {
		if trees[c.x][c.y] <= trees[c.x][i] && i != c.y {
			right++
			break
		}

		if trees[c.x][c.y] > trees[c.x][i] {
			right++
		}
	}

	return top * bottom * right * left
}

func parseInput(fp string) [][]int {
	b, _ := os.ReadFile(fp)
	split := strings.Split(strings.TrimSpace(string(b)), "\n")

	res := make([][]int, len(split))
	for i := 0; i < len(split); i++ {
		res[i] = make([]int, len(split[i]))

		for j, r := range split[i] {
			res[i][j] = int(r - '0')
		}
	}

	return res
}
