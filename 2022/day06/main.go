package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := parseInput("input.txt")

	fmt.Println("part1:", findStartOfPacket(input, 4))
	fmt.Println("part2:", findStartOfPacket(input, 14))
}

func parseInput(fp string) []byte {
	b, _ := os.ReadFile(fp)
	return bytes.TrimSpace(b)
}

func findStartOfPacket(bs []byte, n int) int {
	for i := n; i < len(bs); i++ {
		if isUnique(bs[i-n:i], n) {
			return i
		}
	}

	return -1
}

func isUnique(bs []byte, n int) bool {
	m := make(map[byte]struct{}, n)
	for _, b := range bs {
		if _, ok := m[b]; ok {
			return false
		} else {
			m[b] = struct{}{}
		}
	}

	return true
}
