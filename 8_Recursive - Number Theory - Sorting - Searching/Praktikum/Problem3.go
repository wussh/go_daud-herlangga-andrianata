package main

import "fmt"

var cekprime bool

func main() {
	primesegiempat(5, 2, 5)
}
func primesegiempat(high, wide, start int) {
	arr := []int{}
	cekprime = true
	total := 0

	for i := start; i < 999; i++ {
		cekprime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				cekprime = false
			}
		}
		if cekprime == true {
			arr = append(arr, i)
		}

		if len(arr) == wide*high {
			break
		}
	}

	for i := 1; i <= high*wide; i++ {
		fmt.Print(arr[i-1], ",")
		total += arr[i-1]
		if i%wide == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("total = ", total)
}
