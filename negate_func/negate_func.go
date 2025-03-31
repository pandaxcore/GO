package main

import (
	"fmt"
)

func negate(myBoolean *bool) bool {
	!myBoolean
	return *myBoolean
}

func main() {
	var truth bool = true
	negate(&truth)
	fmt.Println(truth)

	var lies bool = false
	negate(&lies)
	fmt.Println(lies)
}
