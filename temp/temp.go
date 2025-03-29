package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 3; x++ {
		fmt.Println("continue")
		continue
		fmt.Println("after continue")
	}
}