package main

import "fmt"

func main() {
	var z, x int
	fmt.Scan(&z, &x)
	for i := x; i != 0; i-- {
		if z%10 != 0 {
			z--
		} else {
			z = z / 10
		}
	}
	fmt.Println(z)
}
