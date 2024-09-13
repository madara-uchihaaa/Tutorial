package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Gopher")
	greeter("Golang")
}

func greeter(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		// Its used to let thread to sleep for 100 millisecond and then print the value
		fmt.Println("Hello:->", s,i)
	}
}
