package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readInput("input.txt")

	p1 := lowestNumberHash(input, "00000")
	fmt.Println(p1)

	p2 := lowestNumberHash(input, "000000")
	fmt.Println(p2)
}

func readInput(fn string) string {
	contents, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

func lowestNumberHash(key string, prefix string) int {
	var i int

	for {
		data := []byte(fmt.Sprintf("%s%d", key, i))
		hash := md5.Sum(data)

		s := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(s, prefix) {
			return i
		}

		i++
	}
}

// DOESN'T WORK :(
// can't figure out how to only return 1 value from goroutine
func lowestNumberHash2(key string, prefix string) int {
	first := make(chan int, 1)

	for i := 0; i < 100000000; i++ {
		go func(n int) {
			data := []byte(fmt.Sprintf("%s%d", key, n))
			hash := md5.Sum(data)

			s := fmt.Sprintf("%x", hash)

			if strings.HasPrefix(s, prefix) {
				first <- n
			}
		}(i)
	}

	return <-first
}
