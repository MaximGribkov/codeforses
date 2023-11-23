package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for a++; len(map[int]bool{
		a % 10:       true,
		a / 10 % 10:  true,
		a / 100 % 10: true,
		a / 1000:     true}) != 4; a++ {
	}
	fmt.Println(a)
}
