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
	target := rand.Intn(11)
	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You got", 10-guesses, "left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Your guess was HIGH")
		} else {
			fmt.Println("You GUESS is", guess)
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you did not guess my number...")
	}

}
