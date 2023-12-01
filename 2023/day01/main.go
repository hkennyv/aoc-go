package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	input := parseInput("input.txt")

	p1 := part1(input)
	fmt.Println("part1:", p1)

	p2 := part2(input)
	fmt.Println("part2:", p2)
}

func part1(lines []string) int {
	total := 0

	for _, l := range lines {
		val := 0
		i := 0

		// left
		for {
			c := l[i]
			n, err := strconv.Atoi(string(c))
			if err != nil {
				i++
			} else {
				val += 10 * n
				break
			}
		}

		// right
		i = 0
		for {
			c := l[len(l)-1-i]
			n, err := strconv.Atoi(string(c))
			if err != nil {
				i++
			} else {
				val += n
				break
			}
		}

		total += val
	}

	return total
}

func part2(s []string) int {
	total := 0

	for _, l := range s {
		val := 0
		found := false
		i := 0

		// left
		for {
			ss := l[i:]

			n, err := strconv.Atoi(string(ss[0]))
			if err != nil {
				for num := range Nums {
					if strings.HasPrefix(ss, num) {
						val += 10 * Nums[num]
						found = true
						break
					}
				}
			} else {
				val += 10 * n
				break
			}

			if found {
				break
			}

			i++
		}

		// right
		i = 0
		found = false
		for {
			ss := l[:len(l)-i]

			n, err := strconv.Atoi(string(ss[len(ss)-1]))
			if err != nil {
				for num := range Nums {
					if strings.HasSuffix(ss, num) {
						val += Nums[num]
						found = true
						break
					}
				}
			} else {
				val += n
				break
			}

			if found {
				break
			}

			i++
		}

		total += val
	}

	return total
}

func parseInput(fn string) []string {
	b, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(b)), "\n")
}
