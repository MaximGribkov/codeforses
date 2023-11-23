package main

import "fmt"

func main() {
	q, w, e := 0, 0, 0
	var z, x, v, a int
	fmt.Scan(&a)
	for i := a; i != 0; i-- {
		fmt.Scan(&z, &x, &v)
		q += z
		w += x
		e += v
	}
	if q == 0 && w == 0 && e == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
