package main

import "fmt"

func main() {
	// Maps in Golang

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "TypeScript"
	languages["CS"] = "C#"

	fmt.Println(languages)

	// Get value from map

	fmt.Println(languages["JS"])

	// Delete value from map

	delete(languages, "RB")

	fmt.Println(languages)

	// Check if key exists in map

	value, exists := languages["RB"]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exist")
	}

	// Iterate over map
	for key, value := range languages {
		fmt.Println(key, ":", value)
		fmt.Printf("%v : %v\n", key, value)
	}
}
