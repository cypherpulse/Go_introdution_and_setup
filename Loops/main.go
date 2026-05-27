package main

import (
	"fmt"
)

func main() {
	//loops are used to execute a block of code repeatedly until a specified condition is met. In Go, there is only one type of loop, the for loop, which can be used in various ways to achieve different looping behaviors.
	// x:= 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }
	for i := 0; i < 5; i++ {
		fmt.Println("value of x is ", i)
	}
}
