package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winners map[int]struct{}
	numbers map[int]struct{}
}

func main() {
	cards := parseInput("input.txt")

	p1 := part1(cards)
	fmt.Println("part1:", p1)

	p2 := part2(cards)
	fmt.Println("part2:", p2)
}

func part1(cards map[int]Card) int {
	total := 0

	for _, card := range cards {
		count := 0
		for n := range card.numbers {
			if _, ok := card.winners[n]; ok {
				count++
			}
		}

		if count > 0 {
			total += int(math.Pow(2, float64(count-1)))
		}
	}

	return total
}

func part2(cards map[int]Card) int {
	count := make(map[int]int, 0)

	for i := 1; i <= len(cards); i++ {
		count[i]++
		card := cards[i]

		nWinners := 0
		for n := range card.numbers {
			if _, ok := card.winners[n]; ok {
				nWinners++
			}
		}

		for j := 1; j <= nWinners; j++ {
			count[i+j] += count[i]
		}
	}

	total := 0

	for _, v := range count {
		total += v
	}

	return total
}

func parseInput(fn string) map[int]Card {
	b, _ := os.ReadFile(fn)

	cards := make(map[int]Card, 0)

	split := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, s := range split {
		card := parseCard(s)
		cards[card.id] = card
	}

	return cards
}

func parseCard(s string) Card {
	s = s[5:]

	split := strings.Split(s, ":")
	id, _ := strconv.Atoi(strings.TrimSpace(split[0]))

	split = strings.Split(split[1], " | ")

	winners := make(map[int]struct{}, 0)
	for _, ss := range strings.Fields(split[0]) {
		n, _ := strconv.Atoi(ss)
		winners[n] = struct{}{}
	}

	numbers := make(map[int]struct{}, 0)
	for _, ss := range strings.Fields(split[1]) {
		n, _ := strconv.Atoi(ss)
		numbers[n] = struct{}{}
	}

	return Card{id, winners, numbers}
}
