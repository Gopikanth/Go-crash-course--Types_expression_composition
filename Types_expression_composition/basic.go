package main

import "log"

var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "giri"

	log.Println(myString)

	myString = "string"

	var myBool = true
	myBool = false

	log.Println(myBool)

}
