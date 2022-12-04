package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")

	count := 0
	for _, p := range input {
		if fullyOverlaps(p) {
			count++
		}
	}
	fmt.Println("part1:", count)

	count = 0
	for _, p := range input {
		if partialOverlaps(p) {
			count++
		}
	}
	fmt.Println("part2:", count)
}

func parseInput(fp string) [][]int {
	b, _ := os.ReadFile(fp)
	s := strings.Split(strings.TrimSpace(string(b)), "\n")

	pairs := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		pairs[i] = parseLine(s[i])
	}

	return pairs
}

func parseLine(s string) []int {
	res := make([]int, 4)

	split := strings.Split(s, ",")
	left := strings.Split(split[0], "-")
	right := strings.Split(split[1], "-")

	res[0], _ = strconv.Atoi(left[0])
	res[1], _ = strconv.Atoi(left[1])
	res[2], _ = strconv.Atoi(right[0])
	res[3], _ = strconv.Atoi(right[1])

	return res
}

func fullyOverlaps(pairs []int) bool {
	return (pairs[0] <= pairs[2] && pairs[1] >= pairs[3]) || (pairs[2] <= pairs[0] && pairs[3] >= pairs[1])
}

func partialOverlaps(pairs []int) bool {
	return pairs[0] <= pairs[3] && pairs[2] <= pairs[1]
}
