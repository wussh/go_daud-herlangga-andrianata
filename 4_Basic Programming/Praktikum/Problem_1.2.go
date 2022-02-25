package main

import (
	"fmt"
)

func main() {
	var t, r float64
	fmt.Println("Tinggi")
	fmt.Scan(&t)
	fmt.Println("Radius")
	fmt.Scan(&r)
	var pi = 3.14
	var x = 2*pi*r*r + 2*pi*r*t
	fmt.Println(x)
}
