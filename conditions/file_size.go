package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./conditions/my.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print(fileInfo)
}
