package main

import (
	"fmt"
	"sort"
)

func maximumBuyProduct(money int, arr []int) {
	x := 0
	result := 0
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		x = x + arr[i]
		if x > money {
			break
		} else {
			result += 1
		}
	}

	fmt.Println(result)
}

func main() {
	maximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	maximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	maximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	maximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	maximumBuyProduct(0, []int{10000, 30000})
}
