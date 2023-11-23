package main

import (
	"fmt"
)

func main() {
	var a string
	c := "hello"
	count := 0
	fmt.Scan(&a)
	for i := range a {
		if a[i] == c[count] {
			count++
		}
		if count == 5 {
			fmt.Println("YES")
			break
		}
	}
	if count != 5 {
		fmt.Println("NO")
	}
}
