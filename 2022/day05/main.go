package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cheated, didn't bother parsing input from file, just wrote by hand...
var crates = [][]string{
	{"H", "C", "R"},
	{"B", "J", "H", "L", "S", "F"},
	{"R", "M", "D", "H", "J", "T", "Q"},
	{"S", "G", "R", "H", "Z", "B", "J"},
	{"R", "P", "F", "Z", "T", "D", "C", "B"},
	{"T", "H", "C", "G"},
	{"S", "N", "V", "Z", "B", "P", "W", "L"},
	{"R", "J", "Q", "G", "C"},
	{"L", "D", "T", "R", "H", "P", "F", "S"},
}

func main() {
	input := parseInput("input.txt")

	// copy input so we don't mutate the original
	tmp := make([][]string, len(crates))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = make([]string, len(crates[i]))
		copy(tmp[i], crates[i])
	}
	p1(tmp, input)

	for i := 0; i < len(tmp); i++ {
		tmp[i] = make([]string, len(crates[i]))
		copy(tmp[i], crates[i])
	}
	p2(tmp, input)
}

func p1(crates [][]string, input [][]int) {
	for _, l := range input {
		crates[l[0]], crates[l[1]] = moveOneByOne(crates[l[0]], crates[l[1]], l[2])
	}

	fmt.Printf("part1: ")
	for i := 0; i < len(crates); i++ {
		fmt.Printf(crates[i][len(crates[i])-1])
	}
	fmt.Println()
}

func p2(crates [][]string, input [][]int) {
	for _, l := range input {
		crates[l[0]], crates[l[1]] = moveMultiple(crates[l[0]], crates[l[1]], l[2])
	}

	fmt.Printf("part2: ")
	for i := 0; i < len(crates); i++ {
		fmt.Printf(crates[i][len(crates[i])-1])
	}
	fmt.Println()
}

func parseInput(fp string) [][]int {
	b, _ := os.ReadFile(fp)
	s := strings.Split(strings.TrimSpace(string(b)), "\n")[10:]
	res := make([][]int, len(s))

	for i := 0; i < len(s); i++ {
		split := strings.Split(s[i], " ")
		n, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		// instructions are 1-indexed, so subtract 1 to get the crate number
		res[i] = []int{from - 1, to - 1, n}
	}

	return res
}

func moveOneByOne(from, to []string, n int) ([]string, []string) {
	for i := 0; i < n; i++ {
		to = append(to, from[len(from)-1])
		from = from[:len(from)-1]
	}

	return from, to
}

func moveMultiple(from, to []string, n int) ([]string, []string) {
	to = append(to, from[len(from)-n:]...)
	from = from[:len(from)-n]

	return from, to
}
