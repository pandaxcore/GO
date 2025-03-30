package main

import "fmt"

func main() {
	var myInt int = 4
	var myIntPointer = &myInt
	fmt.Println(myInt)
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
