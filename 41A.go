package main

import (
	"fmt"
)

func main() {
	var one, two, b string
	fmt.Scan(&one)
	fmt.Scan(&two)
	for i := len(one); i > 0; i-- {
		b += string(one[i-1])
	}
	if two == b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
