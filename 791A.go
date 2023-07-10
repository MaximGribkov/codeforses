package main

import (
	"fmt"
)

func main() {
	var a, b int
	count := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	for {
		a = a * 3
		b = b * 2
		count++
		if a > b {
			break
		}
	}
	fmt.Println(count)

}
