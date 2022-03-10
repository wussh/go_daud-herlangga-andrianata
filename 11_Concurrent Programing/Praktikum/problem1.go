package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main execution started at time", time.Since(start))

	go getChars("lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do euismod tempor incididunt ut labore et dolore magna aliqua")

	time.Sleep(1900 * time.Millisecond)

	fmt.Println("\nmain execution stopped at time", time.Since(start))
}
