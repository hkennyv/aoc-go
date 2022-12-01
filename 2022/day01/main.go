package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")

	split := strings.Split(input, "\n\n")
	calories := make([]int, 0)

	for _, elf := range split {
		calories = append(calories, countCalories(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	p1 := calories[0]
	fmt.Println("part1:", p1)

	p2 := calories[0] + calories[1] + calories[2]
	fmt.Println("part2:", p2)
}

func parseInput(fp string) string {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(b))
}

func countCalories(s string) int {
	sum := 0
	for _, food := range strings.Split(s, "\n") {
		calories, err := strconv.Atoi(food)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += calories
	}
	return sum
}
