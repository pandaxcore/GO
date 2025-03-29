package main

import (
	"fmt"
)

func pantNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	pantNeeded(4.2, 3.0)
	pantNeeded(5.2, 3.5)
	pantNeeded(5.0, 3.3)
}
