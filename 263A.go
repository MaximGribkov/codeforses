package main

import (
	"fmt"
	"math"
)

func main() {
	var mark string
	index := 0
	for i := 1; i != 26; i++ {
		fmt.Scan(&mark)
		if mark == "1" {
			index = i
		}
	}
	line := math.Ceil(float64(index) / 5)
	column := index % 10
	if column < 10 && column > 5 {
		column = column - 5
	}
	if column == 0 {
		column = 5
	}
	answer := math.Abs(3-line) + math.Abs(3-float64(column))
	fmt.Println(answer)

}
