package main

import (
	"fmt"
)

// func negate(myBoolean *bool) bool {
// 	*myBoolean = !*myBoolean
// 	return *myBoolean
// }

// func main() {
// 	var truth bool = true
// 	negate(&truth)
// 	fmt.Println(truth)

// 	var lies bool = false
// 	negate(&lies)
// 	fmt.Println(lies)
// }

func swapFunction(a *int, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}

func main() {
	var a, b int = 3, 7
	fmt.Printf("Before Swap: a = %d, b = %d\n", a, b)
	swapFunction(&a, &b)
	fmt.Printf("After Swap: a = %d, b = %d\n", a, b)
}
