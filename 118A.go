package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	a = strings.ToLower(a)
	for _, i := range a {
		if i == 97 || i == 111 || i == 121 || i == 101 || i == 117 || i == 105 {
			continue
		} else {
			fmt.Print("." + string(i))
		}
	}
}
