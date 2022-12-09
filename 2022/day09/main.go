package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y int
}

type Pos Vector

var Directions map[string]Vector = map[string]Vector{
	"U": {0, 1},
	"D": {0, -1},
	"R": {1, 0},
	"L": {-1, 0},
}

func main() {
	input := parseInput("input.txt")
	fmt.Println("part1:", simulate(input, 2))
	fmt.Println("part2:", simulate(input, 10))
}

func simulate(input []string, n int) int {
	history := make(map[Pos]struct{}, 0)
	rope := make([]Pos, n)

	history[rope[len(rope)-1]] = struct{}{}

	for _, line := range input {
		d, n := getDirection(line)
		dir := Directions[d]

		for i := 0; i < n; i++ {

			// move current knot
			rope[0].x += dir.x
			rope[0].y += dir.y

			// for each of the next n-1 knots, calculate the vector to the
			// current knot and move accordingly
			//
			// - lagging knot should never move if dist < 2 (because it is
			//   adjacent to leading knot)
			// - if dist is >= 2, move in the unit x and y direction of the
			//   leading knot
			for j := 1; j < len(rope); j++ {
				v, dist := calculateVectorDist(rope[j-1], rope[j])
				if dist >= 2 {
					if v.x != 0 {
						rope[j].x += v.x / int(math.Abs(float64(v.x)))
					}
					if v.y != 0 {
						rope[j].y += v.y / int(math.Abs(float64(v.y)))
					}
				}
			}

			// after updating all ropes, mark tail
			history[rope[len(rope)-1]] = struct{}{}
		}

	}

	return len(history)
}

func calculateVectorDist(a, b Pos) (Vector, float64) {
	return Vector{a.x - b.x, a.y - b.y}, math.Sqrt(math.Pow(float64((a.x-b.x)), 2) + math.Pow(float64(a.y-b.y), 2))
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

func getDirection(s string) (string, int) {
	split := strings.Split(s, " ")
	n, _ := strconv.Atoi(split[1])

	return split[0], n
}
