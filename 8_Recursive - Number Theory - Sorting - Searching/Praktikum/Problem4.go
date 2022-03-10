package main

import "fmt"

func maxSequence(arr []int) int {
	const MaxInt = int(^uint(0) >> 1)
	sum := -MaxInt - 1
	end := 0
	for i := 0; i < len(arr); i++ {
		end = max_end + arr[i]
		if sum < end {
			sum = end
		}

		if end < 0 {
			end = 0
		}

	}

	return sum
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
