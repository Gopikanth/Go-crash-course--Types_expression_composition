package main

import (
	"fmt"
)


func main() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)
}


func changeValueOfPointer(num *int) {
	*num = 25
}