package main

import (
	"errors"
	"fmt"
)

func divide(divident float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return divident / divisor, nil
}

func main() {
	division, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f", division)
	}
}
