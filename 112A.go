package main

import (
	"fmt"
	"strings"
)

func main() {
	var one, two string
	fmt.Scan(&one, &two)
	one = strings.ToLower(one)
	two = strings.ToLower(two)
	if one > two {
		fmt.Println(1)
	} else if one < two {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}
}
