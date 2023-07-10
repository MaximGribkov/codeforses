package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, count int
	var mark string
	fmt.Scan(&n)
	for i := 0; i != n; i++ {
		fmt.Scan(&mark)
		if strings.Count(mark, "+")%2 == 0 && strings.Count(mark, "+") > 1 {
			count++
		}
		if strings.Count(mark, "-")%2 == 0 && strings.Count(mark, "-") > 1 {
			count--
		}
	}
	fmt.Println(count)
}
