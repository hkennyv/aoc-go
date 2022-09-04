package main

import (
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	x, y int
}

func main() {
	input := readInput("input.txt")

	p1 := calculateHouses(input)
	fmt.Println(p1)

	p2 := calculateHousesRobo(input)
	fmt.Println(p2)
}

func readInput(s string) string {
	contents, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

func calculateHouses(s string) int {
	grid := make(map[Coordinate]int, 0)
	current := Coordinate{x: 0, y: 0}

	grid[current] = 1

	for _, ch := range s {
		switch ch {
		case '^':
			current.y++
		case 'v':
			current.y--
		case '>':
			current.x++
		case '<':
			current.x--
		}

		_, ok := grid[current]
		if ok {
			grid[current]++
		} else {
			grid[current] = 1
		}
	}

	return len(grid)
}

func calculateHousesRobo(s string) int {
	grid := make(map[Coordinate]int, 0)

	santa := Coordinate{x: 0, y: 0}
	robosanta := Coordinate{x: 0, y: 0}

	grid[santa] = 2

	current := &santa

	for i, ch := range s {

		if i%2 == 0 {
			current = &santa
		} else {
			current = &robosanta
		}

		switch ch {
		case '^':
			current.y++
		case 'v':
			current.y--
		case '>':
			current.x++
		case '<':
			current.x--
		}

		_, ok := grid[*current]
		if ok {
			grid[*current]++
		} else {
			grid[*current] = 1
		}
	}

	return len(grid)
}
