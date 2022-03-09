package main

import "fmt"

func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
