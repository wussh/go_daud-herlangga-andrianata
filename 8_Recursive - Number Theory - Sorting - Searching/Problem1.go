package main

import (
	"fmt"
)

var memo map[int]int = map[int]int{}

// func getprime(num int) bool {
// 	for x := 2; x < num; x++ {
// 		if num%x == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
func primeX(num int) int {
	var result int
	if result, exist := memo[num]; exist {
		return result
	}
	if num < 2 {
		return 0

	}
	if num%num == 0 {
		return result

	}
	memo[num] = result
	return result
}

func main() {
	fmt.Println(primeX(1))
}
