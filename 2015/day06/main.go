package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands := readInput("input.txt")

	grid := newGrid(1000)

	for _, c := range commands {
		grid.runCommand(c, true)
	}

	p1 := grid.countLights()
	fmt.Println(p1)

	grid = newGrid(1000)

	for _, c := range commands {
		grid.runCommand(c, false)
	}

	p2 := grid.sumLights()
	fmt.Println(p2)
}

func readInput(fn string) []Command {
	contents, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(contents), "\n")
	commands := make([]Command, 0)

	for _, input := range inputs {
		commands = append(commands, inputToCommand(input))
	}

	return commands
}

func inputToCommand(s string) Command {
	var op string
	var c1, c2 Coordinate

	operations := []string{
		"toggle",
		"turn off",
		"turn on",
	}

	for _, p := range operations {
		if strings.HasPrefix(s, p) {
			op = p
			s = strings.TrimPrefix(s, op+" ")
		}
	}

	if op == "" {
		fmt.Println("invalid input")
		os.Exit(1)
	}

	split := strings.Split(s, " ")

	pair := strings.Split(split[0], ",")
	x, _ := strconv.Atoi(pair[0])
	y, _ := strconv.Atoi(pair[1])
	c1 = Coordinate{x: x, y: y}

	pair = strings.Split(split[2], ",")
	x, _ = strconv.Atoi(pair[0])
	y, _ = strconv.Atoi(pair[1])
	c2 = Coordinate{x: x, y: y}

	return Command{
		Operation: op,
		Start:     c1,
		End:       c2,
	}
}
