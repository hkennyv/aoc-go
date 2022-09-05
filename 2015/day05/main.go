package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
var forbiddenDoubles = []string{"ab", "cd", "pq", "xy"}

func main() {
	input := readInput("input.txt")

	rules := []func(string) bool{
		hasThreeVowels,
		containsDoubleLetter,
		doesNotContainPairs,
	}
	p1 := countNiceStrings(input, rules)
	fmt.Println(p1)

	rules = []func(string) bool{
		hasNonOverlappingPair,
		hasRepeatingLetterBetween,
	}
	p2 := countNiceStrings(input, rules)
	fmt.Println(p2)
}

func readInput(fn string) []string {
	contents, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(contents), "\n")
}

func countNiceStrings(ss []string, rules []func(string) bool) int {
	var count int

	for _, s := range ss {
		if isNiceString(s, rules) {
			fmt.Println(s)
			count++
		}
	}

	return count
}

func isNiceString(s string, rules []func(string) bool) bool {
	for _, fn := range rules {
		if !fn(s) {
			return false
		}
	}

	return true
}

func hasThreeVowels(s string) bool {
	var i int

	for _, v := range vowels {
		for _, ch := range s {
			if v == ch {
				i++
			}
		}
	}

	return i >= 3
}

func containsDoubleLetter(s string) bool {
	if len(s) < 2 {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}

	return false
}

func doesNotContainPairs(s string) bool {
	for i := 0; i < len(s); i++ {
		ss := s[i:]
		for _, pair := range forbiddenDoubles {
			if strings.HasPrefix(ss, pair) {
				return false
			}
		}
	}

	return true
}

func hasNonOverlappingPair(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		pair := s[i : i+2]

		for j := range s {
			if j != i-1 && j != i && j != i+1 && strings.HasPrefix(s[j:], pair) {
				return true
			}
		}
	}

	return false
}

func hasRepeatingLetterBetween(s string) bool {
	for i := 0; i <= len(s)-3; i++ {
		triple := s[i : i+3]
		if triple[0] == triple[2] {
			return true
		}
	}

	return false
}
