package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Files in Golangs
	content := "Hello world this is a test file for golangs files package"
	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	l, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(l, "bytes written successfully")

	f, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
	defer file.Close()
}
