package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("input.txt")
	floor, basement := calculateFloor(input)
	fmt.Println(floor)
	fmt.Println(basement)
}

func calculateFloor(input string) (int, int) {
	var floor, basement int
	var set bool

	for i, ch := range input {
		if ch == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 && !set {
			set = true
			basement = i + 1
		}
	}

	return floor, basement
}

func readFile(fn string) string {
	contents, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}
