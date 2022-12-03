package main

import (
	"errors"
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

		m1 := make(map[rune]int)
		m2 := make(map[rune]int)

		fillMap(m1, l)
		fillMap(m2, r)

		c, _ := findCommonN(m1, m2)
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
			ms[j] = make(map[rune]int, 0)
			fillMap(ms[j], b)
		}

		c, _ := findCommonN(ms...)

		keys := make([]rune, 0)
		for k := range c {
			keys = append(keys, k)
		}

		k := keys[0]
		sum += calculatePriority(k)
	}

	return sum
}

func findCommonN(ms ...map[rune]int) (map[rune]int, error) {
	if len(ms) == 1 {
		return nil, errors.New("need more than 1 map")
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

	return common, nil
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

func fillMap(m map[rune]int, s string) {
	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 1
		}
	}
}

func splitCompartments(s string) []string {
	r := make([]string, 0)

	r = append(r, s[0:len(s)/2])
	r = append(r, s[len(s)/2:])

	return r
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}
