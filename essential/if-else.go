// package main

// import "fmt"

// func main() {
// 	if 15%2 == 0 {
// 		fmt.Println("it is even")
// 	} else {
// 		fmt.Println("it is odd")
// 	}
// }

// package main
// import "fmt"
// func main() {
// 	number :=("enter a numbers:"){
// 		if number%2==0{
// 			fmt.Println("it is even numbers")
// 		}
// 		else{
// 			fmt.Println("it is odd numbers")
// 		}
// 	}
// }

// package main
// import "fmt"
// func main(){
// 	name := [] int{2,4,5,6,7,89,78,34,56}{
// 		if name%2==0;
// 		fmt.Println(name,"it is even numbers")
// 	}else{
// 		fmt.Println(name,"it is odd numbers")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	name := 23
// 	if name > 20 {
// 		fmt.Println("it good")
// 		if name > 23 {
// 			fmt.Println("it is not good")
// 		}
// 	} else {
// 		fmt.Println("it is bot same")
// 	}
// }

// package main
// import "fmt"
// func main(){
// 	var a int =1000
// 	if a < 1200{
// 		fmt.Println("it is less than value")
// 	}else{
// 	fmt.println("it is greater than value")

// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a = 10
// 	c := 20
// 	var account_number = 65437
// 	var n string = "siva"
// 	var (
// 		name           = "siva"
// 		version        = 2.4
// 		playstore, key = "apps", "games"
// 	)
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("this is ram")
// 	}
// 	if 15%2 == 0 {
// 		fmt.Println("it is even numbers")
// 	} else {
// 		fmt.Println("it is odd numbers")
// 	}
// 	fmt.Println(a, c, account_number, n, name, version, playstore, key)
// }

// package main

// import "fmt"

// func main() {
// 	numbersArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	channel := make(chan int)

// 	go sumNumbers(numbersArray[:len(numbersArray)/2], channel)
// 	go sumNumbers(numbersArray[len(numbersArray)/2:], channel)

// 	number1 := <-channel
// 	number2 := <-channel

// 	fmt.Println("Sum of two numbers in first goroutine: ", number1)
// 	fmt.Println("Sum of two numbers in second goroutine: ", number2)

// 	fmt.Println("Total : ", number1+number2)
// }

// func sumNumbers(numbers []int, channel chan int) {
// 	result := 0
// 	for _, value := range numbers {
// 		result += value
// 	}
// 	channel <- result
// }

// package main

// func lowerBound(arr[]int){
// 	low :=0
// 	high :=len(arr) - 1
// 	mid :=0

// 	for low < high {
// 		mid = (low + high) /2
// 		if arr
// 	}

// }

// package main

// import (
// 	"fmt"
// )

// func oddEven(num int) {
// 	if num%2 == 0 {
// 		fmt.Println(num, "it is even number")
// 	} else {
// 		fmt.Println(num, "it is odd numbers")
// 	}
// }

// func main() {
// 	oddEven(34)
// 	oddEven(77)
// 	oddEven(7)
// 	oddEven(44)
// 	oddEven(6)
// }

// package main

// func lowerBound(arr[]int){
// 	low :=0
// 	high :=len(arr) - 1
// 	mid :=0

// 	for low < high {
// 		mid = (low + high) /2
// 		if arr[mid]
// // 	}

// }

package main

import "fmt"

func hello(a, b, c int) int {
	sum := a + b + c
	return sum
}
func main() {
	fmt.Println(sum
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
	, hello(2,3))

}
