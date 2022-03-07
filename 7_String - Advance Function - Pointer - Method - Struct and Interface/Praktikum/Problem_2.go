package main

import (
	"fmt"
)

func caesar(offset int, input string) string {

	const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
	const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	offset %= 26
	newText := []byte(input)

	for index, byteValue := range newText {
		if byteValue >= 'a' && byteValue <= 'z' {
			newText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+offset)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			newText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+offset)%26]
		}
	}
	return string(newText)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
