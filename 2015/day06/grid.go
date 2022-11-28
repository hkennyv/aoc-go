package main

import (
	"fmt"
	"os"
)

type Grid struct {
	grid [][]int
	size int
}

type Coordinate struct {
	x, y int
}

type Command struct {
	Operation string
	Start     Coordinate
	End       Coordinate
}

func newGrid(size int) Grid {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	return Grid{grid: grid, size: size}
}

func (g *Grid) runCommand(c Command, binary bool) {
	// lower left corner
	p1 := Coordinate{x: min(c.Start.x, c.End.x), y: min(c.Start.y, c.End.y)}
	// upper right corner
	p4 := Coordinate{x: max(c.Start.x, c.End.x), y: max(c.Start.y, c.End.y)}

	for i := p1.x; i <= p4.x; i++ {
		for j := p1.y; j <= p4.y; j++ {
			if binary {
				switch c.Operation {
				case "toggle":
					g.grid[i][j] = 1 - g.grid[i][j]
				case "turn on":
					g.grid[i][j] = 1
				case "turn off":
					g.grid[i][j] = 0
				default:
					fmt.Println("invalid command", c)
					os.Exit(1)
				}
			} else {
				switch c.Operation {
				case "toggle":
					g.grid[i][j] = g.grid[i][j] + 2
				case "turn on":
					g.grid[i][j] = g.grid[i][j] + 1
				case "turn off":
					g.grid[i][j] = max(0, g.grid[i][j]-1)
				default:
					fmt.Println("invalid command", c)
					os.Exit(1)
				}
			}
		}
	}

}

func (g *Grid) countLights() int {
	count := 0

	for i := range g.grid {
		for j := range g.grid {
			if g.grid[i][j] > 0 {
				count++
			}
		}
	}

	return count
}

func (g *Grid) sumLights() int {
	sum := 0

	for i := range g.grid {
		for j := range g.grid {
			sum += g.grid[i][j]
		}
	}

	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
