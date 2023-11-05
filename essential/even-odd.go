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
------------------------------------------

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("even number", i)
		} else {
			fmt.Println("odd number", i)
		}
	}
}

--------------------------------------------
package main
import (
	"fmt"
)

func main(){
	fmt.Println("enter a number")
	var number int
	fmt.scanln(&number)
if (number%2==0){
	fmt.Println(number,"enter even number")
}else{
	fmt.Println(number,"enter odd number")
}
}

-------------------------------------------

package main
import (
	"fmt"
)
func main(){
	var number =[] int {56,78,44,23,55,23}
	for _,nubs:= range number{
		if nubs %2==o{
			fmt.Println(nubs,"this is odd number")
		}else{
			fmt.Println(nubs,"this is even number")
		}
	}
}


