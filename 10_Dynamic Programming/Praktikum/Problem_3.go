package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Frog(jumps []int) int {
	temp1 := 0
	temp2 := 0
	const MaxInt = int(^uint(0) >> 1)

	for i := 1; i < len(jumps); i++ {
		fv := temp1 + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		fs := MaxInt
		if i > 1 {
			fs = temp2 + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		}

		minValue := min(fv, fs)
		temp2 = temp1
		temp1 = minValue
	}

	return temp1
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
