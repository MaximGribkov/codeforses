package main

import (
	"fmt"
	"strings"
)

func main() {
	var z string
	fmt.Scan(&z)
	if strings.Contains(z, "0000000") || strings.Contains(z, "1111111") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
