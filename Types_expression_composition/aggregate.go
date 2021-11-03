package main

import "fmt"

type Car struct {
	NumberOfTires int
	Seats         bool
	Make          string
	Model         string
	Year          int
}

func main() {

	var myStrings [3]string
	myStrings[0] = "dog"
	myStrings[1] = "fish"
	myStrings[2] = "cat"
	fmt.Println("The first item in the array is", myStrings[0])

	myCar := Car{
		NumberOfTires: 4,
		Seats:         true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
	fmt.Println()
}
