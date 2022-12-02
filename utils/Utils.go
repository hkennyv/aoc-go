package utils

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
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
