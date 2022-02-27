package main

import "fmt"

func cetaktabelperkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d	", i*j)
		}
		println()
	}
}

func main() {
	cetaktabelperkalian(9)
}
