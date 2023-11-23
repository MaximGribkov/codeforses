package main

import (
	"fmt"
)

func main() {
	var in, out, count int
	fmt.Scan(&count)
	c := 0
	max := 0
	for i := count; i > 0; i-- {
		fmt.Scan(&in)
		fmt.Scan(&out)
		c = c - in + out
		if c > max {
			max = c
		}
	}
	fmt.Println(max)
}
