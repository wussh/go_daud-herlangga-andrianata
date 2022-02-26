package main

import "fmt"

func isPrime(num int) bool {
	for x := 2; x < num; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(13))
	fmt.Println(isPrime(17))
	fmt.Println(isPrime(20))
	fmt.Println(isPrime(35))
}
