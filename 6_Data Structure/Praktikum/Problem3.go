package main

import "fmt"

func PairSum(arr []int, target int) {
	mapping := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if mapping[target-arr[i]] == 0 {
			mapping[arr[i]] = i
		} else {
			fmt.Printf("Pair for given sum is (%d, %d).\n", arr[mapping[target-arr[i]]], arr[i])
			return
		}
	}
	fmt.Println("Pair not found in given array.")
}

func main() {
	PairSum([]int{1, 2, 3, 4, 6}, 6)
	PairSum([]int{2, 5, 9, 11}, 11)
	PairSum([]int{1, 3, 5, 7}, 12)
	PairSum([]int{1, 4, 6, 8, 10}, 10)
	PairSum([]int{1, 5, 6, 7}, 6)
}
