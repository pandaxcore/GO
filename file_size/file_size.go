package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./file_size/my.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print(fileInfo.Size(), "\n")
}
