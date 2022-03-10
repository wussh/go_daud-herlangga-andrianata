package main

import "fmt"

func getMinMax(n ...*int) (min int, max int) {
	min = *n[0]
	max = *n[0]
	for _, v := range n {
		if *v < min {
			min = *v
		}
		if *v > max {
			max = *v
		}
	}

	return min, max
}

func main() {
	var n1, n2, n3, n4, n5, n6, min, max int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	fmt.Scan(&n5)
	fmt.Scan(&n6)

	min, max = getMinMax(&n1, &n2, &n3, &n4, &n5, &n6)
	fmt.Println("nilai min =", min)
	fmt.Println("nilai max =", max)
}
