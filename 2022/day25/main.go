package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var SnafuDigits = map[byte]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func main() {
	b, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSpace(string(b)), "\n")

	sum := 0
	for _, line := range input {
		sum += fromSnafu(line)
	}

	fmt.Println("part1:", toSnafu(sum))
}

func fromSnafu(s string) int {
	var res int

	for i := range s {
		b := s[len(s)-1-i]
		v := SnafuDigits[b]

		res += v * int(math.Pow(5, float64(i)))
	}
	return res
}

func toSnafu(n int) string {
	var s, res string

	for n > 0 {
		r := n % 5
		n = n / 5

		s += strconv.Itoa(r)
	}

	carry := false
	for i := 0; i < len(s); i++ {
		c := s[i]

		if carry {
			c += 1
			carry = false
		}

		switch c {
		case '0', '1', '2':
			res += string(c)
		case '3':
			carry = true
			res += "="
		case '4':
			carry = true
			res += "-"
		case '5':
			carry = true
			res += "0"
		}
	}

	if carry {
		res += "1"
	}

	res = reverseString(res)

	return res
}

func reverseString(s string) string {
	var builder strings.Builder
	for i := range s {
		builder.WriteByte(s[len(s)-1-i])
	}
	return builder.String()
}
