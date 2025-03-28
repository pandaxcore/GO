package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess: ")
	input, err := reader.ReadString('\n')
	var guess int
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err = strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Your guess was HIGH")
	}
	fmt.Println(target)
}
