package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://lco.dev"
func main() {
	// Web requests in Golang
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", res)
	fmt.Println("Response is: ", res)
	defer res.Body.Close() // IMP: Close the response body after the function ends

	// Read the response body
	dbytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dbytes))
}
