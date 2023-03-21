package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("siva")
		if i == 3 {
			break
		}
	}

	numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 34, 56}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("sri ram")
	}

	for {
		fmt.Println("siva")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i = 1 + i
	}
}
