package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

var matches map[string][][]int = map[string][][]int{
	"A X": {{3, 1}, {0, 3}},
	"A Y": {{6, 2}, {3, 1}},
	"A Z": {{0, 3}, {6, 2}},
	"B X": {{0, 1}, {0, 1}},
	"B Y": {{3, 2}, {3, 2}},
	"B Z": {{6, 3}, {6, 3}},
	"C X": {{6, 1}, {0, 2}},
	"C Y": {{0, 2}, {3, 3}},
	"C Z": {{3, 3}, {6, 1}},
}

func main() {
	rounds := parseInput("input.txt")

	score := 0
	for _, r := range rounds {
		score += utils.Sum(matches[r][0])
	}

	fmt.Println("part1:", score)

	score = 0
	for _, r := range rounds {
		score += utils.Sum(matches[r][1])
	}

	fmt.Println("part2:", score)
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}
