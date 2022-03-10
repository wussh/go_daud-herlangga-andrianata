package main

import (
	"fmt"
	"strconv"
)

func findMinAndMax(arr []int) string {
	min := arr[0]
	minIndex := 0
	max := arr[0]
	maxIndex := 0
	for i, v := range arr {
		if v < min {
			min = v
			minIndex = i
		}
		if v > max {
			max = v
			maxIndex = i
		}
	}

	return "Min : " + strconv.Itoa(min) + ", Index : " + strconv.Itoa(minIdx) + "\nMax : " + strconv.Itoa(max) + ", Index : " + strconv.Itoa(maxIdx)
}

func main() {
	fmt.Println(findMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println()
	fmt.Println(findMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println()
	fmt.Println(findMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println()
	fmt.Println(findMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println()
	fmt.Println(findMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
	fmt.Println()
}
