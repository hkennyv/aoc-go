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
	buf := make([]byte, n)

	// prefill buf
	for i := 0; i < n; i++ {
		buf[i] = bs[i]
	}

	// first check
	if isUnique(buf, n) {
		return n
	}

	// check rest
	for i := n; i < len(bs); i++ {
		for j := 0; j < n; j++ {
			buf[j] = bs[i-(n-1)+j]
		}

		if isUnique(buf, n) {
			return i + 1
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
