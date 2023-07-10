package main

import (
	"fmt"
)

func main() {
	var l, count int
	var a string
	fmt.Scan(&l, &a)
	for i := 0; i < (l - 1); i++ {
		if a[i] == a[i+1] {
			count++
		}
	}
	fmt.Println(count)

}
