package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	s := make(map[string]bool)
	for _, i := range a {
		s[string(i)] = true
	}
	if len(s)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}

}
