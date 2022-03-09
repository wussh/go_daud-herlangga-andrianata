package main

import (
	"fmt"
	"math/big"
)

func pow(x, n int) *big.Int {
	tmp := big.NewInt(int64(x))
	res := big.NewInt(1)
	for n > 0 {
		temp := new(big.Int)
		if n%2 == 1 {
			temp.Mul(res, tmp)
			res = temp
		}
		temp = new(big.Int)
		temp.Mul(tmp, tmp)
		tmp = temp
		n /= 2
	}
	return res
}

func main() {
	fmt.Println(pow(2, 2))
}
