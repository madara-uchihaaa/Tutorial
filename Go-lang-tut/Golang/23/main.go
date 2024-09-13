package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // WaitGroup is used to wait for the program to finish goroutines. Generally they are pointers

var signals []string

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://flipkart.com",
		"http://myntra.com",
	}
	for _, link := range links {
		go getStatusCode(link)
		wg.Add(1) // Add 1 to the WaitGroup counter
	}
	wg.Wait() // Wait blocks until the WaitGroup counter is zero

	fmt.Println("Total links:", (signals))
}

func getStatusCode(endPoint string) {
	defer wg.Done() // Done decrements the WaitGroup counter by one
	resp, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("Error:", err, "for", endPoint)
	} else {
		signals = append(signals, endPoint)

		fmt.Println("Status Code:", resp.StatusCode, "for", endPoint)
	}
}
