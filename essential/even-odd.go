package main

import (
	"fmt"
)

func oddEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "it is even number")
	} else {
		fmt.Println(num, "it is odd numbers")
	}
}

func main() {
	oddEven(34)
	oddEven(77)
	oddEven(7)
	oddEven(44)
	oddEven(6)
}
