package main

import "fmt"

var cekprime bool

func primesegiempat(high, wide, start int) {
	arr := []int{}
	cekprime = true
	total := 0

	for i := start; i < 999; i++ {
		cekprime = true
		for j := 2; j < i; j++ {
			cekprime = false
		}
	}
	if cekprime == true {
		arr = append(arr, start)
	}
	if len(arr) == wide*high {
	}

	for i := 1; i <= high*wide; i++ {
		fmt.Print(arr[i-1], ",")
		total += arr[i-1]
		if i%wide == 0 {
			fmt.Print("")
		}
	}
	fmt.Println("total = ", total)
}

func main() {
	primesegiempat(5, 2, 5)
}
