package main

import "fmt"

func main() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

random:
	fmt.Println("Random")

	for i, day := range days {
		println(i, day)
	}

	for _, day := range days {
		println(day)
	}

	for i := 0; i < len(days); i++ {
		println(days[i])
	}

	i := 0
	for i < len(days) {
		println(days[i])
		i++
		goto random
	}

}
