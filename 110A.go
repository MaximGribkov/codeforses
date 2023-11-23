package main

import (
	"fmt"
	"strings"
)

func main() {
	var z string
	fmt.Scan(&z)
	count := strings.Count(z, "7") + strings.Count(z, "4")
	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
