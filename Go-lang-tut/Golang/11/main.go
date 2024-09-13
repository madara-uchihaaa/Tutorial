package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
	writeTable(67)
	num, val := processData([]int{1, 2, 3, 4, 5})
	fmt.Println(num, val)
}
func writeTable(a int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}

func processData(val []int) (int, string) {
	return 1, "hello"
}
