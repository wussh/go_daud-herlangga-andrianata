package main

import (
	"fmt"
)

func pangkat(base, Pangkat int) int {
	output := 1
	for Pangkat != 0 {
		output *= base
		Pangkat -= 1
	}
	fmt.Printf("%d", output)
	return 0
}
func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}

// func main() {
// 	var pangkat, base int
// 	fmt.Print("Enter Base:")
// 	fmt.Scanln(&base)
// 	fmt.Print("Enter pangkat:")
// 	fmt.Scanln(&pangkat)

// 	output := 1
// 	for pangkat != 0 {
// 		output *= base
// 		pangkat -= 1
// 	}
// 	fmt.Printf("The Output of power calculation is %d", output)
// }
