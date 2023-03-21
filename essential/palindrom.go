// package main

// import "fmt"

// func main() {
// 	var palndrom int
// 	fmt.Println("enter the values plindrom:")
// 	fmt.Scanln(&palndrom)
// 	reverse := 0

// 	if palndrom == reverse {
// 		fmt.Println(palndrom, "it is palindrom crt")
// 	} else {
// 		fmt.Println(palndrom, "it is not palindrom value")
// 	}
// }

package main

import "fmt"

func palindrom(num int) string {
	name := num
	var pall int
	call := o

	for num > 0 {
		pall = num % 10
		call = (call * 10) + pall
		num = num / 10
	}
	if name == call {
		return "it is palindrom"
	} else {
		return "it is not a palindrom"
	}
}

func main() {
	fmt.Println(131)
	fmt.Println(141)
}
