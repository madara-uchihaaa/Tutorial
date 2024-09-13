package main

import "fmt"

func main() {
	defer fmt.Println("Hello world with defer 1");
	defer fmt.Println("Hello world with defer 2");
	fmt.Println("Hello world");
}