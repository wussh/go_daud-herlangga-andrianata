package main

import "fmt"

func simpleequation(a, b, c int) {
	for x := 1; x <= a; x++ {
		for y := 1; y <= a; y++ {
			for z := 1; z <= a; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	simpleequation(1, 2, 3)
	simpleequation(6, 6, 14)
}
