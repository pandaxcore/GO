// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Print("Enter a grade: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	input = strings.TrimSpace(input)
// 	var grade float64
// 	grade, err = strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var status string
// 	if grade >= 60 {
// 		status = "passing"
// 	} else {
// 		status = "failing"
// 	}
// 	fmt.Println("A grade of", grade, "is", status)

// }

package main

import (
	"example/go/src/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
