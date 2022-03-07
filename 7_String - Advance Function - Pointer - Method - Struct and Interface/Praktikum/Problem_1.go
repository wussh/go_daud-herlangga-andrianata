package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	if len(a) < len(b) {
		if strings.Contains(b, a) {
			return a
		}
	} else {
		if strings.Contains(a, b) {
			return b
		}
	}

	return "tidak ada yang sama"
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILA", "ILALANG"))
}
