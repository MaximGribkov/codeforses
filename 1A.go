package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d", int(math.Ceil(a/c)*math.Ceil(b/c)))
}
