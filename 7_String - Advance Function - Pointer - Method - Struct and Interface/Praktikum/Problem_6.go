package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Decode() string
	Encode() string
}

func (data *student) Decode() string {
	const normalChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const codeChar = "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	newText := []byte(data.nameEncode)

	for index, byteValue := range newText {
		if byteValue >= 'a' && byteValue <= 'z' {
			lowerNormalChar := strings.ToLower(normalChar)
			lowerCoderChar := strings.ToLower(codeChar)
			for j := 0; j < 26; j++ {
				if byteValue == lowerNormalChar[j] {
					newText[index] = lowerCoderChar[j]
					break
				}
				if byteValue < 'a' || byteValue > 'z' {
					newText[index] = byteValue
					break
				}
			}
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			for j := 0; j < 26; j++ {
				if byteValue == normalChar[j] {
					newText[index] = codeChar[j]
					break
				}
				if byteValue < 'A' || byteValue > 'Z' {
					newText[index] = byteValue
					break
				}
			}
		}

	}
	return string(newText)
}

func (data *student) Encode() string {
	const normalChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const codeChar = "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	newText := []byte(data.name)

	for index, byteValue := range newText {
		if byteValue >= 'a' && byteValue <= 'z' {
			lowerNormalChar := strings.ToLower(normalChar)
			lowerCoderChar := strings.ToLower(codeChar)
			for j := 0; j < 26; j++ {
				if byteValue == lowerNormalChar[j] {
					newText[index] = lowerCoderChar[j]
					break
				}
				if byteValue < 'a' || byteValue > 'z' {
					newText[index] = byteValue
					break
				}
			}
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			for j := 0; j < 26; j++ {
				if byteValue == normalChar[j] {
					newText[index] = codeChar[j]
					break
				}
				if byteValue < 'A' || byteValue > 'Z' {
					newText[index] = byteValue
					break
				}
			}
		}

	}
	return string(newText)
}

func main() {
	var menu int
	var s = student{}
	var c Chiper = &s

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput student's name : ")
		fmt.Scan(&s.name)
		fmt.Println("\nEncode student's name", s.name, "is :", c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput student's encode name : ")
		fmt.Scan(&s.nameEncode)
		fmt.Println("\nDecode student's name", s.nameEncode, "is :", c.Decode())
	} else {
		fmt.Println("Wrong input")
	}
}
