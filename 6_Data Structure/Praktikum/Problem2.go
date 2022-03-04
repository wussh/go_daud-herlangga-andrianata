package main

import (
	"fmt"
	"strconv"
)

func main() {
	var foundMap = map[byte]int{}
	data := "76523752"
	for i := 0; i < len(data); i++ {
		character := data[i]
		_, exist := foundMap[character]
		if !exist {
			foundMap[character] = 0
		}
		foundMap[character] += 1
		fmt.Println("i", i, "character", character, "foundMap", foundMap)
	}
	for character, found := range foundMap {
		fmt.Println("character", character, "found", found)
		if found == 1 {
			angkaStr := fmt.Sprintf("%c", character)
			angka, _ := strconv.Atoi(angkaStr)
			fmt.Println(angka)
			// fmt.Printf("%c\n", character)
		}
	}

}
