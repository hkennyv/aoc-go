package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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

// this is my best attempt at a shitty concurrent solution in go. this one uses
// a worker pool pattern to spawn a pool of 50 workers that will receive a
// number and calculate the hash. if it finds a match, it passes it to a
// channel to be returned
//
// it performs a little better than the non-concurrent solution @ 2.875s vs
// 3.566s for the concurrent on my machine
//
// it still has some bugs and if we increase the size of the buffered input
// channel, it crashes because it will find >1 matches and try to send them
// to the found channel
func lowestNumberHash2(key string, prefix string) int {
	var wg sync.WaitGroup

	numWorkers := 50

	input := make(chan int, 100000)
	found := make(chan int)

	stop := make(chan interface{})

	// spawn numWorkers workers who will run until they receive a stop signal
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(key, prefix, input, found, stop)
		}()
	}

	return receiver(numWorkers, input, found, stop)
}

// this worker consumes input from the input channel indefinitely until a stop
// signal is sent. once it finds a hash with matching prefix, it sends the
// result over the found channel
func worker(key string, prefix string, input <-chan int, found chan<- int, stop <-chan interface{}) {
	for {
		select {
		case i := <-input:
			data := []byte(fmt.Sprintf("%s%d", key, i))
			hash := md5.Sum(data)

			s := fmt.Sprintf("%x", hash)

			if strings.HasPrefix(s, prefix) {
				found <- i
			}
		case <-stop:
			return
		}
	}
}

// the receiver will check the found channel for any results and continue to
// increment and put numbers into the input channel for the workers to consume
// once it receives a value from the found channel, it sends the stop signals
// and returns the result
func receiver(numWorkers int, input chan<- int, found <-chan int, stop chan<- interface{}) int {
	var i int

	for {
		select {
		case res := <-found:
			// send stop to all workers, this can still produce a race
			// condition though if the workers try to send another result to
			// the found channel before they get stopped
			for i := 0; i < numWorkers; i++ {
				stop <- struct{}{}
			}

			return res
		default:
			input <- i
		}
		i++
	}
}
