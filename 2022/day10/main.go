package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycleCheck map[int]struct{} = map[int]struct{}{
	20:  {},
	60:  {},
	100: {},
	140: {},
	180: {},
	220: {},
}

func main() {
	s := readFile("input.txt")
	ops := makeOps(s)
	fmt.Println(run(ops))
}

func run(ops []int) int {
	sum := 0
	x := 1

	for i := 0; i < len(ops); i++ {
		pos := i % 40
		cycle := i + 1

		if pos == 0 {
			fmt.Printf("\n")
		}
		if pos >= x-1 && pos <= x+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if _, ok := cycleCheck[cycle]; ok {
			sum += x * (cycle)
		}

		x += ops[i]
	}
	fmt.Println()
	return sum
}

func readFile(fp string) string {
	b, _ := os.ReadFile(fp)
	return string(b)
}

func makeOps(s string) []int {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	ops := make([]int, 0)

	for _, l := range lines {
		if l == "noop" {
			ops = append(ops, 0)
		} else {
			// insert a "noop" before each addx instruction
			ops = append(ops, 0, parseOp(l))
		}
	}
	return ops
}

func parseOp(s string) int {
	split := strings.Split(s, " ")
	n, _ := strconv.Atoi(split[1])
	return n
}
