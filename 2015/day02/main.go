package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

type Present struct {
	sides []int
}

func main() {
	input := readFile("input.txt")

	p1 := totalWrappingPaper(input)
	fmt.Println(p1)

	p2 := totalRibbon(input)
	fmt.Println(p2)
}

func totalWrappingPaper(presents []Present) int {
	papers := make([]int, 0)
	for _, present := range presents {
		papers = append(papers, present.CalculateSquareFeet())
	}
	return utils.TotalIntSlice(papers)
}

func totalRibbon(presents []Present) int {
	ribbons := make([]int, 0)
	for _, present := range presents {
		ribbons = append(ribbons, present.CalcualteRibbon())
	}
	return utils.TotalIntSlice(ribbons)
}

func readFile(fn string) []Present {
	contents, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")
	presents := make([]Present, 0)

	for _, line := range lines {
		presents = append(presents, MakePresent(line))
	}

	return presents
}

func MakePresent(s string) Present {
	split := strings.Split(s, "x")
	sides := make([]int, 0)

	for _, side := range split {
		num, err := strconv.Atoi(side)
		if err != nil {
			log.Fatal(err)
		}

		sides = append(sides, num)
	}

	sort.Ints(sides)

	present := Present{sides: sides}

	return present
}

func (p Present) CalculateSquareFeet() int {
	var res int
	sidesArea := []int{p.sides[0] * p.sides[1], p.sides[1] * p.sides[2], p.sides[2] * p.sides[0]}

	for _, area := range sidesArea {
		res += 2 * area
	}

	res += utils.MinIntSlice(sidesArea)

	return res
}

func (p Present) CalcualteRibbon() int {
	res := 2*p.sides[0] + 2*p.sides[1]
	res += p.sides[0] * p.sides[1] * p.sides[2]
	return res
}
