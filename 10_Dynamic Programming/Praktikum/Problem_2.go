package main

import "fmt"

func fibo(n int) int {
	lol := make([]int, n+1, n+2)
	if n < 2 {
		lol = lol[0:2]
	}
	lol[0] = 0
	lol[1] = 1
	for i := 2; i <= n; i++ {
		lol[i] = lol[i-1] + lol[i-2]
	}
	return lol[n]
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(8))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))

}
