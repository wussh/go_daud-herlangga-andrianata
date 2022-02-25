package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	for i := 1; i <= d; i++ {
		if d%i == 0 {
			fmt.Println(i)
		}
	}
}
