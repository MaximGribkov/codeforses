package main

import (
	"fmt"
	"math"
)

func main() {
	var n, h, o, s float64
	var count int
	fmt.Scan(&n, &h)
	for i := n; i > 0; i-- {
		fmt.Scan(&o)
		s = o / h
		count = count + int(math.Ceil(s))
	}
	fmt.Println(count)
}
