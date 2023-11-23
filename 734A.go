package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	var z string
	fmt.Scan(&a)
	fmt.Scan(&z)
	if strings.Count(z, "A") > strings.Count(z, "D") {
		fmt.Println("Anton")
	} else if strings.Count(z, "D") > strings.Count(z, "A") {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
