package utils

import "fmt"

func HelloWorld() {
	fmt.Println("HELLO WORLD!")
}

func MinIntSlice(nums []int) int {
	var min int

	for i, num := range nums {
		if i == 0 || num < min {
			min = num
		}
	}

	return min
}

func TotalIntSlice(nums []int) int {
	var res int
	for _, num := range nums {
		res += num
	}

	return res
}
