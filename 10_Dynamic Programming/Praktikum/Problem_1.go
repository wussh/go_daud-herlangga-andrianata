package main

import "fmt"

var memo map[int]int = map[int]int{}

func fibomemo(n int) int {
	var result int
	if result, exist := memo[n]; exist {
		return result
	}
	//Tambahan if
	if n == 0 {
		return 0
	}
	if n <= 1 {
		return 1
	} else {
		result = fibomemo(n-1) + (fibomemo(n - 2))
	}
	memo[n] = result
	return result
}

func main() {
	fmt.Println(fibomemo(0))
	fmt.Println(fibomemo(1))
	fmt.Println(fibomemo(2))
	fmt.Println(fibomemo(3))
	fmt.Println(fibomemo(4))
	fmt.Println(fibomemo(5))
	fmt.Println(fibomemo(6))
	fmt.Println(fibomemo(7))
	fmt.Println(fibomemo(8))
	fmt.Println(fibomemo(9))
	fmt.Println(fibomemo(10))

}
