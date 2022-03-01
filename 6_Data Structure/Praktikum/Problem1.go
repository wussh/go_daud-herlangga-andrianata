package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	c := append(arrayA, arrayB...)
	fmt.Println(c)
	return c
}

func main() {
	//test case
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "brian"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"jin", "steve", "brian"}))

	fmt.Println(ArrayMerge([]string{"hwoaramg"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
