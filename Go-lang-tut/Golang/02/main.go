package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// User Input

	welcome := "Welcome to my awesome program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok syntax || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("You typed: ", input)
	fmt.Printf("You typed: %T", input);
}
