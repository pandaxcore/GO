package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func double(number *int) {
	*number *= 2
}

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	*myIntPointer = 42
	fmt.Println(myInt)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var amount int = 6
	double(&amount)
	fmt.Println(amount)
}
