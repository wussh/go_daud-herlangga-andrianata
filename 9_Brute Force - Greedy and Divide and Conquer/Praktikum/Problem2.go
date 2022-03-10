package main

import (
	"fmt"
	"sort"
)

/*
   Go program
   Greedy algorithm to find minimum number of coins
*/

// Find minimum coins whose sum make a given value
func minNoOfCoins(coins []int, n int, value int) {
	if value <= 0 {
		return
	}
	// Sort the given coins
	sort.Ints(coins)
	// Use to collect coins
	var record = make([]int, 0)
	// Auxiliary variables
	var sum int = 0
	var i int = n - 1
	var c int = 0
	//  Find the change coins by given value
	for i >= 0 && sum < value {
		// Get coin
		c = coins[i]
		for c+sum <= value {
			// Add coin
			record = append(record, c)
			// Update sum
			sum += c
		}
		// Reduce position of element
		i--
	}
	fmt.Println("\n Given value is ", value)
	if sum == value {
		// When change are possible.
		// Display resultant coins.
		for v := 0; v < len(record); v++ {
			fmt.Print("  ", record[v])
		}
	} else {
		fmt.Println(" Full change are not possible")
	}
}
func main() {

	var coins = []int{
		10000,
		5000,
		2000,
		1000,
		500,
		200,
		100,
		20,
		10,
		1,
	}
	var n int = len(coins)
	// Value = 23
	minNoOfCoins(coins, n, 23)
	// Value = 38
	minNoOfCoins(coins, n, 38)
	// Value = 7
	minNoOfCoins(coins, n, 7)
}
