package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var a int
	var word string
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Panicf("error scan, err: %s", err)
	}
	for i := 0; i != a; i++ {
		_, err = fmt.Fscan(os.Stdin, &word)
		if err != nil {
			log.Printf("error scan in for, err: %s", err)
		}
		if len(word) > 10 {
			res := fmt.Sprintf("%s%s%s", string(word[0]), strconv.Itoa(len(word)-2), string(word[len(word)-1]))
			fmt.Println(res)
		} else {
			fmt.Println(word)
		}
	}
}
