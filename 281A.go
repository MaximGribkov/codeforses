package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(strings.ToUpper(string(a[0])) + a[1:])
}
