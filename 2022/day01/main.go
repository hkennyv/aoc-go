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
	fm := makeFoodMap(input)

	p1 := findNElvesWithMostFood(fm, 1)
	fmt.Println("part1:", p1)

	p2 := findNElvesWithMostFood(fm, 3)
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

func makeFoodMap(s string) map[int]int {
	fm := make(map[int]int)
	elves := strings.Split(s, "\n\n")

	for i, elf := range elves {
		fm[i] = countCalories(elf)
	}

	return fm
}

func countCalories(s string) int {
	sum := 0
	for _, food := range strings.Split(s, "\n") {
		calories, err := strconv.Atoi(strings.TrimSpace(food))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += calories
	}
	return sum
}

func findNElvesWithMostFood(elves map[int]int, n int) int {
	sum := 0

	values := make([]int, 0)

	for _, calories := range elves {
		values = append(values, calories)
	}

	// sorts in ascending order
	sort.Ints(values)

	// sum top n
	for i := 0; i < n; i++ {
		sum += values[len(values)-1-i]
	}

	return sum
}
