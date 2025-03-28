package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println(tax)

	var total float32 = float32(price) + float32(tax)
	fmt.Println(total)

	var availableFunds int = 120
	fmt.Println(availableFunds)
	fmt.Println(availableFunds > int(total))
}
