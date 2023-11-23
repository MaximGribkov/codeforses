package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	var c string
	fmt.Scan(&a, &b, &c)
	for ; b > 0; b-- {
		c = strings.Replace(c, "BG", "GB", -1)
	}
	fmt.Println(c)
}
