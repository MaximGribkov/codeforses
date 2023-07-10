package main

import (
	"fmt"
)

func main() {
	var n, k, mark int
	var ints []int
	count := 0
	fmt.Scan(&n, &k)
	for i := 0; i != n; i++ {
		fmt.Scan(&mark)
		ints = append(ints, mark)
	}
	for _, i := range ints {
		if i >= ints[k-1] && i != 0 {
			count++
		}
	}
	fmt.Println(count)
}
