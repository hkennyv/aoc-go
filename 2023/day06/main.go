package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func (r Race) getPossibilities() int {
	n := 0

	for i := 1; i < r.time; i++ {
		t := r.time - i
		d := i * t

		if d > r.distance {
			n++
		}
	}
	return n
}

func main() {
	races := parseInput("input.txt")
	res := 1
	for _, r := range races {
		res *= r.getPossibilities()
	}
	fmt.Println("part1:", res)

	race := parseInput2("input.txt")
	res = race.getPossibilities()
	fmt.Println("part2:", res)
}

func parseInput(fn string) []Race {
	b, _ := os.ReadFile(fn)

	times := make([]int, 0)
	distances := make([]int, 0)

	split := strings.Split(strings.TrimSpace(string(b)), "\n")

	for _, f := range strings.Fields(split[0])[1:] {
		n, _ := strconv.Atoi(f)
		times = append(times, n)
	}

	for _, f := range strings.Fields(split[1])[1:] {
		n, _ := strconv.Atoi(f)
		distances = append(distances, n)
	}

	races := make([]Race, 0)

	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}

	return races
}

func parseInput2(fn string) Race {
	b, _ := os.ReadFile(fn)

	split := strings.Split(strings.TrimSpace(string(b)), "\n")

	time := strings.Join(strings.Fields(split[0])[1:], "")
	distance := strings.Join(strings.Fields(split[1])[1:], "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	return Race{t, d}
}
