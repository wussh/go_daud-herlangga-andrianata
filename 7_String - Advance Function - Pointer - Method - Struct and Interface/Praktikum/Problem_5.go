package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (data Student) getAvarage() float64 {
	result := 0
	for _, v := range data.score {
		result += v
	}

	return float64(result / len(data.score))
}

func (data Student) getMax() (max int, name string) {
	max = data.score[0]
	name = data.name[0]
	for i, v := range data.score {
		if v > max {
			max = v
			name = data.name[i]
		}
	}

	return max, name
}

func (data Student) getMin() (min int, name string) {
	min = data.score[0]
	name = data.name[0]
	for i, v := range data.score {
		if v < min {
			min = v
			name = data.name[i]
		}
	}

	return min, name
}

func main() {
	var data = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i+1) + " Student's Name : ")
		fmt.Scan(&name)
		data.name = append(data.name, name)

		var score int
		fmt.Print("Input " + strconv.Itoa(i+1) + " Score : ")
		fmt.Scan(&score)
		data.score = append(data.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", data.getAvarage())
	scoreMax, nameMax := data.getMax()
	fmt.Println("Max score student is", nameMax, "(", scoreMax, ")")
	scoreMin, nameMin := data.getMin()
	fmt.Println("Min score student is", nameMin, "(", scoreMin, ")")
}
