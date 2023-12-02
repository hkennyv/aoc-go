package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

// Each game has multiple draws
// Each draw is a mapping of marble color to quantity
type Draw map[string]int
type Game struct {
	gameId int
	draws  []Draw
}

func main() {
	games := parse_input("input.txt")

	p1 := part1(games)
	fmt.Println("part1:", p1)

	p2 := part2(games)
	fmt.Println("part2:", p2)
}

func part1(games []Game) int {
	total := 0
	bag := Draw{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range games {

		impossible := false

		for _, draw := range game.draws {
			for k := range draw {
				if draw[k] > bag[k] && !impossible {
					impossible = true
				}
			}

			if impossible {
				continue
			}
		}

		if !impossible {
			total += game.gameId
		}
	}

	return total
}

func part2(games []Game) int {
	total := 0

	for _, game := range games {
		val := 1
		bag := Draw{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, draw := range game.draws {
			for k := range draw {
				bag[k] = utils.Max(bag[k], draw[k])
			}
		}

		for k := range bag {
			val *= bag[k]
		}

		total += val
	}

	return total
}

func parse_input(fn string) []Game {
	b, _ := os.ReadFile(fn)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	games := make([]Game, 0)

	for _, l := range lines {
		games = append(games, parse_game(l))
	}

	return games
}

func parse_game(s string) Game {
	split := strings.Split(s, ": ")
	gameId, _ := strconv.Atoi(strings.Split(strings.TrimSpace(split[0]), " ")[1])

	split = strings.Split(split[1], "; ")
	draws := make([]Draw, 0)

	for _, s := range split {
		d := make(Draw, 0)
		picks := strings.Split(s, ", ")

		for _, p := range picks {
			ss := strings.Split(p, " ")
			n, _ := strconv.Atoi(ss[0])
			color := ss[1]

			d[color] = n
		}

		draws = append(draws, d)
	}

	return Game{
		gameId,
		draws,
	}
}
