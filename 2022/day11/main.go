package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation func(item int) int
	Test      func(n int) int
	Inspected int
	Divisor   int
}

func (m *Monkey) add(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) play(monkeys Monkeys, lcm bool) {
	for _, item := range m.Items {
		var new int

		if lcm {
			new = m.Operation(item) % monkeys.getLcm()
		} else {
			new = m.Operation(item) / 3
		}

		to := m.Test(new)
		monkeys[to].add(new)
		m.Inspected++
	}
	m.Items = []int{}
}

type Monkeys []*Monkey

func (ms Monkeys) getLcm() int {
	lcm := 1
	for _, m := range ms {
		lcm *= m.Divisor
	}
	return lcm
}

// implement sort.Interface
func (ms Monkeys) Len() int           { return len(ms) }
func (ms Monkeys) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }
func (ms Monkeys) Less(i, j int) bool { return ms[i].Inspected < ms[j].Inspected }

func main() {
	input := parseInput("input.txt")
	monkeys := make(Monkeys, 0)
	for _, l := range input {
		monkeys = append(monkeys, parseMonkey(l))
	}

	// part1 - 20 rounds
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.play(monkeys, false)
		}
	}

	sort.Sort(monkeys)
	fmt.Println("part1:", monkeys[len(monkeys)-1].Inspected*monkeys[len(monkeys)-2].Inspected)

	// part2 - 10000 rounds
	monkeys = make(Monkeys, 0)
	for _, l := range input {
		monkeys = append(monkeys, parseMonkey(l))
	}
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.play(monkeys, true)
		}
	}

	sort.Sort(monkeys)
	fmt.Println("part2:", monkeys[len(monkeys)-1].Inspected*monkeys[len(monkeys)-2].Inspected)
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n\n")
}

func parseMonkey(p string) *Monkey {
	split := strings.Split(p, "\n")

	// parse items
	s := strings.TrimSpace(strings.Split(split[1], ":")[1])
	ss := strings.Split(s, ", ")
	items := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		n, _ := strconv.Atoi(ss[i])
		items[i] = n
	}

	// parse operation
	var op, operand string
	var operation func(item int) int

	s = strings.TrimSpace(strings.Split(split[2], ":")[1])
	fmt.Sscanf(s, "new = old %s %s", &op, &operand)
	operandn, err := strconv.Atoi(operand)
	if err != nil {
		if op == "*" {
			operation = func(item int) int {
				return item * item
			}
		} else {
			operation = func(item int) int {
				return item + item
			}
		}
	} else {
		if op == "*" {
			operation = func(item int) int {
				return item * operandn
			}
		} else {
			operation = func(item int) int {
				return item + operandn
			}
		}

	}

	// parse test
	var divisor, totrue, tofalse int
	fmt.Sscanf(split[3], "  Test: divisible by %d", &divisor)
	fmt.Sscanf(split[4], "    If true: throw to monkey %d", &totrue)
	fmt.Sscanf(split[5], "    If false: throw to monkey %d", &tofalse)

	test := func(n int) int {
		if n%divisor == 0 {
			return totrue
		} else {
			return tofalse
		}
	}

	return &Monkey{
		Items:     items,
		Operation: operation,
		Test:      test,
		Inspected: 0,
		Divisor:   divisor,
	}
}
