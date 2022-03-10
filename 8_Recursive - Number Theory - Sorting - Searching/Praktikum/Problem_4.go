package main

import "fmt"

func maxSequence(arr []int) int {
	const MaxInt = int(^uint(0) >> 1)

	max_sum := -MaxInt - 1
	max_end := 0

	for i := 0; i < len(arr); i++ {
		max_end = max_end + arr[i]
		if max_sum < max_end {
			max_sum = max_end
		}

		if max_end < 0 {
			max_end = 0
		}

	}

	return max_sum
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
