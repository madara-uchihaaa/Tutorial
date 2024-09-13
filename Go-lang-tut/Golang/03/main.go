package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my awesome program")
	fmt.Println("Enter the rating of my program:")

	reader := bufio.NewReader(os.Stdin)

	in, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", in)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(in), 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
