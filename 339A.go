package main

import (
	"fmt"
	"sort"
)

func main() {
	var one, answer string
	fmt.Scan(&one)
	a := make([]int, 0, len(one))
	for _, i := range one {
		if i != 43 {
			a = append(a, int(i))
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, c := range a {
		answer = answer + string(c) + "+"
	}
	fmt.Println(answer[0 : len(answer)-1])
}
