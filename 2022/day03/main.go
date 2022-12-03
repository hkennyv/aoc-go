package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := parseInput("input.txt")

	fmt.Println("part1:", p1(input))
	fmt.Println("part2:", p2(input))
}

func p1(bags []string) int {

	sum := 0

	for _, bag := range bags {
		cs := splitCompartments(bag)
		l, r := cs[0], cs[1]

		m1 := makeRuneMap(l)
		m2 := makeRuneMap(r)

		c := findCommonN(m1, m2)
		keys := make([]rune, 0)
		for k := range c {
			keys = append(keys, k)
		}

		k := keys[0]
		sum += calculatePriority(k)
	}

	return sum
}

func p2(bags []string) int {
	sum := 0

	for i := 0; i < len(bags); i += 3 {
		triplet := bags[i : i+3]

		ms := make([]map[rune]int, len(triplet))

		for j, b := range triplet {
			ms[j] = makeRuneMap(b)
		}

		c := findCommonN(ms...)

		keys := make([]rune, 0)
		for k := range c {
			keys = append(keys, k)
		}

		k := keys[0]
		sum += calculatePriority(k)
	}

	return sum
}

func findCommonN(ms ...map[rune]int) map[rune]int {
	if len(ms) == 1 {
		return ms[0]
	}

	common := make(map[rune]int, 0)

	for _, m := range ms {
		for k := range m {
			common[k] = 1
		}
	}

	for k := range common {
		for _, m := range ms {
			if _, ok := m[k]; !ok {
				delete(common, k)
			}
		}
	}

	return common
}

func calculatePriority(r rune) int {
	if r < 'A' || r > 'z' {
		fmt.Println(r, "is out of range")
		os.Exit(1)
	}

	if r >= 'a' && r <= 'z' {
		return int(r) - 'a' + 1
	} else {
		return int(r) - 'A' + 27
	}
}

func makeRuneMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c] = 1
	}

	return m
}

func fillMap(m map[rune]int, s string) {

	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 1
		}
	}
}

func splitCompartments(s string) []string {
	return []string{
		s[0 : len(s)/2],
		s[len(s)/2:],
	}
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}
