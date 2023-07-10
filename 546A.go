package main

import (
	"fmt"
)

func main() {
	var price, balance, num int
	fmt.Scan(&price, &balance, &num)
	for i := 1; i < num+1; i++ {
		balance = balance - price*i
	}
	if balance < 0 {
		fmt.Println(balance * -1)
	} else {
		fmt.Println(0)
	}

}
