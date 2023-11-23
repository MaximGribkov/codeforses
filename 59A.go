package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var a string
	countOne := 0
	countTwo := 0
	fmt.Scan(&a)
	for _, i := range a {
		if unicode.IsUpper(i) {
			countOne++
		} else {
			countTwo++
		}
	}
	if countOne > countTwo {
		fmt.Println(strings.ToUpper(a))
	} else {
		fmt.Println(strings.ToLower(a))
	}
}
