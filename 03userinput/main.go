package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok syntax || error ok syntax
	// _ stores error
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of Rating is %T ", input)
}
