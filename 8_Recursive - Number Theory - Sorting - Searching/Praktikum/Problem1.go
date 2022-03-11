package main

import (
	"fmt"
)

// var memo map[int]int = map[int]int{}

// // func getprime(num int) bool {
// // 	for x := 2; x < num; x++ {
// // 		if num%x == 0 {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }
// // prime_slice2.go
// //
// // get a slice of primes from 2 to a given value limit
// //
// // info on numeric limits:
// // https://golang.org/ref/spec#Numeric_types
// // int or int32, signed 32-bit integer (-2147483648 to 2147483647)
// // uint64 is unsigned 64-bit integer (0 to 18446744073709551615)
// //
// // online play at:
// //
// //
// // tested with Go version 1.4.2   by vegaseat  5may2015

// // package main

// // import (
// // 	"fmt"
// // )

// // return a slice of primes from 2 to limit (inclusive)
// // uses Sieve of Eratosthenes algorithm
// func primes_dns(limit int) []int {
// 	limit = limit + 1
// 	// creates a slice of false with length of limit
// 	bools := make([]bool, limit)
// 	// implies up to the sqrt of limit
// 	for k := 2; k*k <= limit; k++ {
// 		if bools[k] != true {
// 			for ix := k * k; ix < limit; ix += k {
// 				bools[ix] = true
// 			}
// 		}
// 	}
// 	// index of remaining False in bools = a prime number
// 	primes := []int{}
// 	for ix := 2; ix < limit; ix++ {
// 		if bools[ix] != true {
// 			primes = append(primes, ix)
// 		}
// 	}
// 	return primes
// }

// func main() {
// 	fmt.Println("A slice of prime numbers:")

// 	prime_slice := primes_dns(100)
// 	fmt.Printf("%v\n", prime_slice)
// }

// /*
// A slice of prime numbers:
// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67]
// */

// func isprime(n int) bool {
// 	if n == 1 || n == 0 {
// 		return false
// 	}
// 	for i := 2; i <= n/2; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func primeX(num int) int {
// 	n := 100
// 	for i := 1; 1 <= n; i++ {
// 		if isprime(i) {
// 			fmt.Println(i, " ")
// 		}
// 	}
// 	return -1
// var result int
// if result, exist := memo[num]; exist {
// 	return result
// }
// if num < 2 {
// 	return 0

// }
// if num%num == 0 {
// 	return result

// }
// memo[num] = result
// return result
// }

func main() {
	fmt.Println(primeX(1))

	// var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// for i, element := range a {
	// 	fmt.Printf("%d: %d\n", i, element)
	// }
}

// func fibomemo(n int) int {
// 	var result int
// 	if result, exist := memo[n]; exist {
// 		return result
// 	}
// 	//Tambahan if
// 	if n == 0 {
// 		return 0
// 	}
// 	if n <= 1 {
// 		return 1
// 	} else {
// 		result = fibomemo(n-1) + (fibomemo(n - 2))
// 	}
// 	memo[n] = result
// 	return result
// }

// func main() {
// 	// fmt.Println(fibomemo(0))
// 	// fmt.Println(fibomemo(1))
// 	// fmt.Println(fibomemo(2))
// 	// fmt.Println(fibomemo(3))
// 	// fmt.Println(fibomemo(4))
// 	// fmt.Println(fibomemo(5))
// 	// fmt.Println(fibomemo(6))
// 	// fmt.Println(fibomemo(7))
// 	// fmt.Println(fibomemo(8))
// 	// fmt.Println(fibomemo(9))
// 	fmt.Println(fibomemo(10))

// }

var hasil int
var prime bool

func primeX(number int) int {
	max := 999
	prime = true
	ke := 0
	for i := 2; i <= max; i++ {
		prime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
			}

		}
		if prime == true {

			ke += 1
			hasil = i
		}
		if ke == number {
			break
		}
	}
	return hasil
}
