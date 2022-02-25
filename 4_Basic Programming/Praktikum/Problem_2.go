package main

import (
	"fmt"
)

func main() {
	var studentscore int = 80
	if studentscore > 100 {
		fmt.Println("Nilai Invalid")
	} else if studentscore >= 80 {
		println("A")
	} else if studentscore >= 65 {
		println("B")
	} else if studentscore >= 50 {
		println("C")
	} else if studentscore >= 35 {
		println("D")
	} else if studentscore >= 0 {
		println("E")
	} else if studentscore < 0 {
		println("Nilai Invalid")
	}

}
